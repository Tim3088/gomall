package service

import (
	"Go-Mall/app/checkout/rpc"
	checkout "Go-Mall/rpc_gen/kitex_gen/checkout"
	"Go-Mall/rpc_gen/kitex_gen/order"
	"Go-Mall/rpc_gen/kitex_gen/payment"
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
)

type CheckoutOrderService struct {
	ctx context.Context
} // NewCheckoutOrderService new CheckoutOrderService
func NewCheckoutOrderService(ctx context.Context) *CheckoutOrderService {
	return &CheckoutOrderService{ctx: ctx}
}

// Run create note info
func (s *CheckoutOrderService) Run(req *checkout.CheckoutOrderReq) (resp *checkout.CheckoutOrderResp, err error) {
	// Finish your business logic.
	// 获取订单物品
	var orderItems []*order.OrderItem
	res, err := rpc.OrderClient.GetOrderItems(s.ctx, &order.GetOrderItemsReq{OrderId: req.OrderId})
	if err != nil {
		return nil, err
	}
	orderItems = res.OrderItems

	// 计算总价
	var total float32
	for _, v := range orderItems {
		total += v.Cost
	}

	// 支付请求
	payReq := &payment.ChargeReq{
		UserId:  req.UserId,
		OrderId: req.OrderId,
		Amount:  total,
		CreditCard: &payment.CreditCardInfo{
			CreditCardNumber:          req.CreditCard.CreditCardNumber,
			CreditCardExpirationMonth: req.CreditCard.CreditCardExpirationMonth,
			CreditCardExpirationYear:  req.CreditCard.CreditCardExpirationYear,
			CreditCardCvv:             req.CreditCard.CreditCardCvv,
		},
	}

	//付款
	paymentResult, err := rpc.PaymentClient.Charge(s.ctx, payReq)
	if err != nil {
		return nil, err
	}
	klog.Info(paymentResult)

	//设置订单为已经支付
	_, err = rpc.OrderClient.MarkOrderPaid(s.ctx, &order.MarkOrderPaidReq{
		OrderId: req.OrderId,
	})
	if err != nil {
		return nil, err
	}

	//构造响应
	resp = &checkout.CheckoutOrderResp{
		TransactionId: paymentResult.TransactionId,
	}
	return
}
