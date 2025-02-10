package serversuite

import (
	"Go-Mall/common/mtl"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/server"
	prometheus "github.com/kitex-contrib/monitor-prometheus"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	registryconsul "github.com/kitex-contrib/registry-consul"
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
		// 设置 Server 侧的 Service 信息，用于服务注册
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: s.CurrentServiceName,
		}),
		// 对当前服务设置 prometheus 监控
		server.WithTracer(prometheus.NewServerTracer(
			"",
			"",
			prometheus.WithDisableServer(true),
			prometheus.WithRegistry(mtl.Registry),
		)),
		// 设置链路追踪
		server.WithSuite(tracing.NewServerSuite()),
	}

	// 创建consul注册器
	r, err := registryconsul.NewConsulRegister(s.RegistryAddr)
	if err != nil {
		klog.Fatal(err)
	}

	// 将consul注册器放入server配置
	opts = append(opts, server.WithRegistry(r))

	return opts
}
