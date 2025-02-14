package utils

import (
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"strings"
)

func GetTokenFromContext(ctx *app.RequestContext) (string, error) {
	// 获取请求头中的Bearer Token
	tokenWithBearer := ctx.GetHeader("Authorization")
	token := strings.TrimPrefix(string(tokenWithBearer), "Bearer ")
	if token == "" {
		return "", errors.New("token not found")
	}
	return token, nil
}
