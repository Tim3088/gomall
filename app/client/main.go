package main

import (
	"Go-Mall/app/client/biz/router"
	"Go-Mall/app/client/conf"
	"Go-Mall/app/client/infra/rpc"
	"Go-Mall/common/mtl"
	"github.com/cloudwego/hertz/pkg/app/server"
	hertzprom "github.com/hertz-contrib/monitor-prometheus"
	"github.com/hertz-contrib/sessions"
	"github.com/hertz-contrib/sessions/redis"
	"github.com/joho/godotenv"
	"os"
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

	registerMiddleware(h)

	router.GeneratedRegister(h)

	h.Spin()
}

func registerMiddleware(h *server.Hertz) {
	// 注册session中间件
	store, err := redis.NewStore(100, "tcp", conf.GetConf().Redis.Address, "", []byte(os.Getenv("SESSION_SECRET")))
	if err != nil {
		panic(err)
	}
	store.Options(sessions.Options{MaxAge: 86400, Path: "/"})
	rs, err := redis.GetRedisStore(store)
	if err == nil {
		rs.SetSerializer(sessions.JSONSerializer{})
	}
	h.Use(sessions.New("Go-Mall", store))
}
