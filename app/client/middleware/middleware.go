package middleware

import (
	"github.com/cloudwego/hertz/pkg/app/server"
)

// RegisterMiddlewares 注册所有中间件
func RegisterMiddlewares(h *server.Hertz) {
	RegisterCasbinMiddleware()
}
