package serversuite

import (
	"Go-Mall/common/mtl"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/config-consul/consul"
	consulServer "github.com/kitex-contrib/config-consul/server"
	prometheus "github.com/kitex-contrib/monitor-prometheus"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	registryconsul "github.com/kitex-contrib/registry-consul"
	"os"
)

// CommonServerSuite 定义自定义Suite
type CommonServerSuite struct {
	CurrentServiceName string
	RegistryAddr       string
}

// Options 实现Suite接口的Options方法
func (s CommonServerSuite) Options() []server.Option {
	// 基础服务配置
	opts := []server.Option{
		server.WithMetaHandler(transmeta.ServerHTTP2Handler),
	}

	// 创建consul注册器
	r, err := registryconsul.NewConsulRegister(s.RegistryAddr)
	if err != nil {
		klog.Fatal(err)
	}

	// 将consul注册器放入server配置
	opts = append(opts, server.WithRegistry(r))

	// 如果开启配置中心，则将consul配置中心放入server配置
	if os.Getenv("CONFIG_CENTER_ENABLED") == "true" {
		consulNodes := os.Getenv("CONFIG_CENTER_NODES")
		if consulNodes != "" {
			// 创建consul客户端
			consulClient, err := consul.NewClient(consul.Options{})
			if err != nil {
				klog.Error(err)
			} else {
				// 将consul配置中心放入server配置
				opts = append(opts, server.WithSuite(consulServer.NewSuite(s.CurrentServiceName, consulClient)))
			}
		}
	}

	opts = append(opts,
		// 设置 Server 侧的 Service 信息，用于服务注册
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: s.CurrentServiceName,
		}),
		// 设置链路追踪
		server.WithSuite(tracing.NewServerSuite()),
		// 对当前服务设置 prometheus 监控
		server.WithTracer(prometheus.NewServerTracer(
			"",
			"",
			prometheus.WithDisableServer(true),
			prometheus.WithRegistry(mtl.Registry),
		)),
		// 设置链路追踪
		server.WithSuite(tracing.NewServerSuite()),
	)

	return opts
}
