package service

import (
	"Go-Mall/app/client/hertz_gen/client/order"
	"Go-Mall/app/client/infra/rpc"
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

func (h *OrderListService) Run(req *order.ListOrderReq) (resp *ListOrderResp, err error) {
	res, err := rpc.OrderClient.ListOrder(h.Context, &rpcorder.ListOrderReq{UserId: req.UserId})
	if err != nil {
		return nil, err
	}
	return &ListOrderResp{res.Orders}, nil
}
