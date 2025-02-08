package utils

import (
	"Go-Mall/app/user/biz/model"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/jwt"
	"os"
	"time"
)

var (
	JwtUtil     *jwt.HertzJWTMiddleware
	IdentityKey = "identity"
)

// InitJwt 初始化 JWT 工具
func InitJwt() {
	var err error
	JwtUtil, err = jwt.New(&jwt.HertzJWTMiddleware{
		Realm:         "test zone",                                        // 认证领域
		Key:           []byte(os.Getenv("JWT_KEY")),                       // 用于签名的密钥
		Timeout:       24 * time.Hour,                                     // 令牌过期时间
		MaxRefresh:    24 * time.Hour,                                     // 令牌最大刷新时间
		TokenLookup:   "header: Authorization, query: token, cookie: jwt", // 令牌查找位置
		TokenHeadName: "Bearer",                                           // 令牌头名称
		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
			// 身份处理函数
			claims := jwt.ExtractClaims(ctx, c)
			return jwt.MapClaims{
				IdentityKey: claims[IdentityKey],
				"role":      "user", // 角色信息 用于casbin
			}
		},
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			// 生成令牌负载
			if v, ok := data.(*model.User); ok {
				return jwt.MapClaims{
					IdentityKey: v.ID,
					"role":      "user", // 角色信息 用于casbin
				}
			}
			return jwt.MapClaims{}
		},
	})
	if err != nil {
		panic(err)
	}
}
