package service

import (
	auth "Go-Mall/rpc_gen/kitex_gen/auth"
	"context"
	"github.com/golang-jwt/jwt/v4"
	"os"
	"time"
)

type DeliverTokenByRPCService struct {
	ctx context.Context
} // NewDeliverTokenByRPCService new DeliverTokenByRPCService
func NewDeliverTokenByRPCService(ctx context.Context) *DeliverTokenByRPCService {
	return &DeliverTokenByRPCService{ctx: ctx}
}

// Run create note info
func (s *DeliverTokenByRPCService) Run(req *auth.DeliverTokenReq) (resp *auth.DeliveryResp, err error) {
	// 生成 JWT 令牌
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": req.UserId,
		"role":    "user",
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	// 签名令牌
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_KEY")))
	if err != nil {
		return nil, err
	}

	// 构建响应
	resp = &auth.DeliveryResp{
		Token: tokenString,
	}

	return resp, nil
}
