package middleware

// RegisterMiddlewares 注册所有中间件
func RegisterMiddlewares() {
	RegisterCasbinMiddleware()
	RegisterLogMiddleware()
}
