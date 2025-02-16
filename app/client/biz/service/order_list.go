package service

import (
	"Go-Mall/app/client/infra/rpc"
	clientutils "Go-Mall/app/client/utils"
	"Go-Mall/rpc_gen/kitex_gen/auth"
	rpcorder "Go-Mall/rpc_gen/kitex_gen/order"
	"context"

	"github.com/cloudwego/hertz/pkg/app"
)

type OrderListService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

type ListOrderResp struct {
	Orders []*rpcorder.Order `json:"orders"`
}

func NewOrderListService(Context context.Context, RequestContext *app.RequestContext) *OrderListService {
	return &OrderListService{RequestContext: RequestContext, Context: Context}
}

func (h *OrderListService) Run() (resp *ListOrderResp, err error) {
	token, err := clientutils.GetTokenFromContext(h.RequestContext)
	if err != nil {
		return nil, err
	}
	// 通过token获取用户信息
	authResp, err := rpc.AuthClient.VerifyTokenByRPC(h.Context, &auth.VerifyTokenReq{Token: token})
	if err != nil {
		return nil, err
	}
	res, err := rpc.OrderClient.ListOrder(h.Context, &rpcorder.ListOrderReq{UserId: uint32(authResp.UserId)})
	if err != nil {
		return nil, err
	}
	return &ListOrderResp{res.Orders}, nil
}
