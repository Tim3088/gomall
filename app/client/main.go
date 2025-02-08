package main

import (
	"Go-Mall/app/client/biz/router"
	"Go-Mall/app/client/biz/utils"
	"Go-Mall/app/client/conf"
	"Go-Mall/app/client/infra/rpc"
	"Go-Mall/app/client/middleware"
	clientutils "Go-Mall/app/client/utils"
	"Go-Mall/common/mtl"
	"github.com/cloudwego/hertz/pkg/app/server"
	hertzprom "github.com/hertz-contrib/monitor-prometheus"
	"github.com/joho/godotenv"
)

var (
	ServiceName  = clientutils.ServiceName
	MetricsPort  = conf.GetConf().Hertz.MetricsPort
	RegistryAddr = conf.GetConf().Hertz.RegistryAddr
)

func main() {
	_ = godotenv.Load()

	// 初始化Jwt
	utils.InitJwt()
	// 初始化prometheus
	mtl.InitMetric(ServiceName, MetricsPort, RegistryAddr)
	rpc.InitClient()

	address := conf.GetConf().Hertz.Address

	// 对当前服务设置 prometheus 监控
	h := server.New(server.WithHostPorts(address), server.WithTracer(
		hertzprom.NewServerTracer(
			"",
			"",
			hertzprom.WithRegistry(mtl.Registry),
			hertzprom.WithDisableServer(true),
		),
	),
	)

	// 注册中间件
	middleware.RegisterMiddlewares(h)

	router.GeneratedRegister(h)

	h.Spin()
}
