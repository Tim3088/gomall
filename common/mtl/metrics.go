package mtl

import (
	"Go-Mall/common/utils"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/pkg/registry"
	"github.com/cloudwego/kitex/server"
	consul "github.com/kitex-contrib/registry-consul"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net"
	"net/http"
)

// Registry prometheus注册器
var Registry *prometheus.Registry

// InitMetric 初始化prometheus
func InitMetric(serviceName string, metricsPort string, registryAddr string) {
	Registry = prometheus.NewRegistry()
	Registry.MustRegister(collectors.NewGoCollector())
	Registry.MustRegister(collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}))

	// 根据consul地址创建consul注册器
	r, _ := consul.NewConsulRegister(registryAddr)

	// 获取本地ip
	localIp := utils.MustGetLocalIPv4()
	ip, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s%s", localIp, metricsPort))
	if err != nil {
		hlog.Error(err)
	}

	// 注册consul服务
	registryInfo := &registry.Info{
		ServiceName: "prometheus",
		Addr:        ip,
		Weight:      1,
		Tags:        map[string]string{"service": serviceName},
	}
	_ = r.Register(registryInfo)

	// 通过server的关闭钩子，在服务关闭时注销consul服务
	server.RegisterShutdownHook(func() {
		r.Deregister(registryInfo)
	})

	// 启动http服务
	http.Handle("/metrics", promhttp.HandlerFor(Registry, promhttp.HandlerOpts{}))
	go http.ListenAndServe(metricsPort, nil)
}
