package service

import (
	"Go-Mall/app/order/biz/dal/mysql"
	"Go-Mall/app/order/biz/model"
	"Go-Mall/rpc_gen/kitex_gen/cart"
	order "Go-Mall/rpc_gen/kitex_gen/order"
	"context"
	"github.com/cloudwego/kitex/pkg/kerrors"
)

type GetOrderItemsService struct {
	ctx context.Context
} // NewGetOrderItemsService new GetOrderItemsService
func NewGetOrderItemsService(ctx context.Context) *GetOrderItemsService {
	return &GetOrderItemsService{ctx: ctx}
}

// Run create note info
func (s *GetOrderItemsService) Run(req *order.GetOrderItemsReq) (resp *order.GetOrderItemsResp, err error) {
	// Finish your business logic.
	o, err := model.GetOrder(s.ctx, mysql.DB, req.OrderId)
	if err != nil {
		return nil, kerrors.NewBizStatusError(50001, err.Error())
	}
	var orderItems []*order.OrderItem
	for _, oi := range o.OrderItems {
		orderItems = append(orderItems, &order.OrderItem{
			Cost: oi.Cost,
			Item: &cart.CartItem{
				ProductId: oi.ProductId,
				Quantity:  oi.Quantity,
			},
		})
	}
	return &order.GetOrderItemsResp{
		OrderItems: orderItems,
	}, nil
}
