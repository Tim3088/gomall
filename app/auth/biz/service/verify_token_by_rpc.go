package service

import (
	auth "Go-Mall/rpc_gen/kitex_gen/auth"
	"context"
	"github.com/golang-jwt/jwt/v4"
	"os"
)

type VerifyTokenByRPCService struct {
	ctx context.Context
} // NewVerifyTokenByRPCService new VerifyTokenByRPCService
func NewVerifyTokenByRPCService(ctx context.Context) *VerifyTokenByRPCService {
	return &VerifyTokenByRPCService{ctx: ctx}
}

// Run create note info
func (s *VerifyTokenByRPCService) Run(req *auth.VerifyTokenReq) (resp *auth.VerifyResp, err error) {
	// Finish your business logic.
	// 通过jwt解析token
	token := req.Token
	// 解析token
	claims := jwt.MapClaims{}
	_, err = jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_KEY")), nil
	})
	if err != nil {
		return nil, err
	}
	// 构建响应
	resp = &auth.VerifyResp{
		UserId: int32(claims["user_id"].(float64)),
		Role:   uint32(claims["role"].(float64)),
	}
	return resp, nil
}
