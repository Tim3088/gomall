package main

import (
	"Go-Mall/app/order/biz/service"
	"Go-Mall/rpc_gen/kitex_gen/order"
	"context"
)

// OrderServiceImpl implements the last service interface defined in the IDL.
type OrderServiceImpl struct{}

// PlaceOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) PlaceOrder(ctx context.Context, req *order.PlaceOrderReq) (resp *order.PlaceOrderResp, err error) {
	resp, err = service.NewPlaceOrderService(ctx).Run(req)

	return resp, err
}

// ListOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) ListOrder(ctx context.Context, req *order.ListOrderReq) (resp *order.ListOrderResp, err error) {
	resp, err = service.NewListOrderService(ctx).Run(req)

	return resp, err
}

// MarkOrderPaid implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) MarkOrderPaid(ctx context.Context, req *order.MarkOrderPaidReq) (resp *order.MarkOrderPaidResp, err error) {
	resp, err = service.NewMarkOrderPaidService(ctx).Run(req)

	return resp, err
}

// GetOrderItems implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) GetOrderItems(ctx context.Context, req *order.GetOrderItemsReq) (resp *order.GetOrderItemsResp, err error) {
	resp, err = service.NewGetOrderItemsService(ctx).Run(req)

	return resp, err
}
