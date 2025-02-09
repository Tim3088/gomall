package service

import (
	"Go-Mall/app/order/biz/dal/mysql"
	"Go-Mall/app/order/biz/model"
	order "Go-Mall/rpc_gen/kitex_gen/order"
	"context"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PlaceOrderService struct {
	ctx context.Context
} // NewPlaceOrderService new PlaceOrderService
func NewPlaceOrderService(ctx context.Context) *PlaceOrderService {
	return &PlaceOrderService{ctx: ctx}
}

// Run create note info
func (s *PlaceOrderService) Run(req *order.PlaceOrderReq) (resp *order.PlaceOrderResp, err error) {
	if len(req.OrderItems) == 0 {
		err = kerrors.NewBizStatusError(50001, "order items is empty")
		return
	}
	err = mysql.DB.Transaction(func(tx *gorm.DB) error {
		orderId, _ := uuid.NewUUID()
		o := &model.Order{
			OrderId:      orderId.String(),
			UserId:       req.UserId,
			UserCurrency: req.UserCurrency,
			Consignee: model.Consignee{
				Email: req.Email,
			},
		}
		if req.Address != nil {
			a := req.Address
			o.Consignee.StreetAddress = a.StreetAddress
			o.Consignee.City = a.City
			o.Consignee.State = a.State
			o.Consignee.Country = a.Country
			o.Consignee.ZipCode = a.ZipCode
		}
		if err := tx.Create(o).Error; err != nil {
			return err
		}
		var items []model.OrderItem
		for _, v := range req.OrderItems {
			items = append(items, model.OrderItem{
				OrderIdRefer: orderId.String(),
				ProductId:    v.Item.ProductId,
				Quantity:     uint32(v.Item.Quantity),
				Cost:         v.Cost,
			})
		}
		if err := tx.Create(items).Error; err != nil {
			return err
		}
		resp = &order.PlaceOrderResp{
			Order: &order.OrderResult{
				OrderId: orderId.String(),
			},
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return
}
