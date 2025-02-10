package service

import (
	"Go-Mall/app/client/infra/rpc"
	rpccart "Go-Mall/rpc_gen/kitex_gen/cart"
	rpcorder "Go-Mall/rpc_gen/kitex_gen/order"
	"context"

	order "Go-Mall/app/client/hertz_gen/client/order"
	"github.com/cloudwego/hertz/pkg/app"
)

type PlaceOrderService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewPlaceOrderService(Context context.Context, RequestContext *app.RequestContext) *PlaceOrderService {
	return &PlaceOrderService{RequestContext: RequestContext, Context: Context}
}

type PlaceOrderResp struct {
	Order *rpcorder.OrderResult `json:"order"`
}

func (h *PlaceOrderService) Run(req *order.PlaceOrderReq) (resp *PlaceOrderResp, err error) {
	res, err := rpc.OrderClient.PlaceOrder(h.Context, &rpcorder.PlaceOrderReq{
		UserId:       req.UserId,
		UserCurrency: req.UserCurrency,
		Address: &rpcorder.Address{
			State:         req.Address.State,
			City:          req.Address.City,
			StreetAddress: req.Address.StreetAddress,
			ZipCode:       req.Address.ZipCode,
			Country:       req.Address.Country,
		},
		Email: req.Email,
		OrderItems: func() []*rpcorder.OrderItem {
			items := make([]*rpcorder.OrderItem, len(req.OrderItems))
			for i, item := range req.OrderItems {
				items[i] = &rpcorder.OrderItem{
					Item: &rpccart.CartItem{
						ProductId: item.Item.ProductId,
						Quantity:  item.Item.Quantity,
					},
					Cost: item.Cost,
				}
			}
			return items
		}(),
	})
	if err != nil {
		return nil, err
	}
	return &PlaceOrderResp{Order: res.Order}, nil
}
