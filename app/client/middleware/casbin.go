package middleware

import (
	"Go-Mall/app/client/infra/rpc"
	rpcauth "Go-Mall/rpc_gen/kitex_gen/auth"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/casbin"
	"log"
	"strings"
)

var CasbinAuth *casbin.Middleware

// RegisterCasbinMiddleware 注册 casbin 中间件
func RegisterCasbinMiddleware() {
	var err error
	CasbinAuth, err = casbin.NewCasbinMiddleware("../../conf/model.conf",
		"../../conf/policy.csv", subjectFromJwt)
	if err != nil {
		log.Fatal(err)
	}
}

// subjectFromJwt 从 jwt 中获取 subject
func subjectFromJwt(ctx context.Context, c *app.RequestContext) string {
	token := c.GetHeader("Authorization")
	// 将token转化为string
	tokenStr := strings.TrimPrefix(string(token), "Bearer ")
	if tokenStr == "" {
		log.Println("Authorization header is empty or invalid")
		return ""
	}
	res, err := rpc.AuthClient.VerifyTokenByRPC(ctx, &rpcauth.VerifyTokenReq{Token: tokenStr})
	if err != nil {
		log.Println(err)
		return ""
	}
	if res.Role == 1 {
		return "admin"
	}
	return "user"
}
