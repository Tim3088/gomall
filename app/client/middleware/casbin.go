package middleware

import (
	"Go-Mall/app/client/biz/utils"
	"Go-Mall/app/client/infra/rpc"
	rpcauth "Go-Mall/rpc_gen/kitex_gen/auth"
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
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
		utils.SendErrResponse(ctx, c, consts.StatusOK, errors.New("Authorization header is empty"))
		return ""
	}
	res, err := rpc.AuthClient.VerifyTokenByRPC(ctx, &rpcauth.VerifyTokenReq{Token: tokenStr})
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, errors.New("verify token failed"))
		return ""
	}
	if res.Role == 1 {
		return "admin"
	}
	return "user"
}
