package main

import (
	"Go-Mall/app/client/biz/router"
	"Go-Mall/app/client/conf"
	"Go-Mall/app/client/infra/rpc"
	"Go-Mall/common/mtl"
	"github.com/cloudwego/hertz/pkg/app/server"
	hertzprom "github.com/hertz-contrib/monitor-prometheus"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	rpc.InitClient()

	address := conf.GetConf().Hertz.Address

	h := server.New(server.WithHostPorts(address), server.WithTracer(
		hertzprom.NewServerTracer(
			"",
			"",
			hertzprom.WithRegistry(mtl.Registry),
			hertzprom.WithDisableServer(true),
		),
	),
	)

	router.GeneratedRegister(h)

	h.Spin()
}
