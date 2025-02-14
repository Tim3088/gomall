package service

import (
	"Go-Mall/app/client/infra/rpc"
	clientutils "Go-Mall/app/client/utils"
	"Go-Mall/rpc_gen/kitex_gen/auth"
	"Go-Mall/rpc_gen/kitex_gen/cart"
	"context"

	common "Go-Mall/app/client/hertz_gen/client/common"
	"github.com/cloudwego/hertz/pkg/app"
)

type EmptyCartService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewEmptyCartService(Context context.Context, RequestContext *app.RequestContext) *EmptyCartService {
	return &EmptyCartService{RequestContext: RequestContext, Context: Context}
}

func (h *EmptyCartService) Run() (resp *common.Empty, err error) {
	token, err := clientutils.GetTokenFromContext(h.RequestContext)
	if err != nil {
		return nil, err
	}
	// 通过token获取用户信息
	authResp, err := rpc.AuthClient.VerifyTokenByRPC(h.Context, &auth.VerifyTokenReq{Token: token})
	if err != nil {
		return nil, err
	}
	_, err = rpc.CartClient.EmptyCart(h.Context, &cart.EmptyCartReq{
		UserId: uint32(authResp.UserId),
	})
	if err != nil {
		return nil, err
	}
	return nil, nil
}
