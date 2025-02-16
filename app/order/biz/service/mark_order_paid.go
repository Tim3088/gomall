package service

import (
	"Go-Mall/app/order/biz/dal/mysql"
	"Go-Mall/app/order/biz/model"
	order "Go-Mall/rpc_gen/kitex_gen/order"
	"context"
)

type MarkOrderPaidService struct {
	ctx context.Context
} // NewMarkOrderPaidService new MarkOrderPaidService
func NewMarkOrderPaidService(ctx context.Context) *MarkOrderPaidService {
	return &MarkOrderPaidService{ctx: ctx}
}

// Run create note info
func (s *MarkOrderPaidService) Run(req *order.MarkOrderPaidReq) (resp *order.MarkOrderPaidResp, err error) {
	// Finish your business logic.
	orderId := req.OrderId
	result := mysql.DB.WithContext(s.ctx).Model(&model.Order{}).Where("order_id = ?", orderId).Update("paid", true)
	return &order.MarkOrderPaidResp{}, result.Error
}
