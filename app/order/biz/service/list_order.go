package service

import (
	"Go-Mall/app/order/biz/dal/mysql"
	"Go-Mall/app/order/biz/model"
	"Go-Mall/rpc_gen/kitex_gen/cart"
	order "Go-Mall/rpc_gen/kitex_gen/order"
	"context"
	"github.com/cloudwego/kitex/pkg/kerrors"
)

type ListOrderService struct {
	ctx context.Context
} // NewListOrderService new ListOrderService
func NewListOrderService(ctx context.Context) *ListOrderService {
	return &ListOrderService{ctx: ctx}
}

// Run create note info
func (s *ListOrderService) Run(req *order.ListOrderReq) (resp *order.ListOrderResp, err error) {
	list, err := model.ListOrder(s.ctx, mysql.DB, req.UserId)
	if err != nil {
		return nil, kerrors.NewBizStatusError(50001, err.Error())
	}
	var orders []*order.Order
	for _, v := range list {
		var items []*order.OrderItem
		for _, oi := range v.OrderItems {
			items = append(items, &order.OrderItem{
				Cost: oi.Cost,
				Item: &cart.CartItem{
					ProductId: oi.ProductId,
					Quantity:  int32(oi.Quantity),
				},
			})
		}
		orders = append(orders, &order.Order{
			OrderId:      v.OrderId,
			UserId:       v.UserId,
			UserCurrency: v.UserCurrency,
			Email:        v.Consignee.Email,
			Address: &order.Address{
				StreetAddress: v.Consignee.StreetAddress,
				City:          v.Consignee.City,
				Country:       v.Consignee.Country,
				State:         v.Consignee.State,
				ZipCode:       v.Consignee.ZipCode,
			},
			OrderItems: items,
		})
	}
	resp = &order.ListOrderResp{
		Orders: orders,
	}
	return
}
