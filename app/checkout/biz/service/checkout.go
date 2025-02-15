package service

import (
	"Go-Mall/app/checkout/rpc"
	"Go-Mall/rpc_gen/kitex_gen/cart"
	checkout "Go-Mall/rpc_gen/kitex_gen/checkout"
	"Go-Mall/rpc_gen/kitex_gen/order"
	"Go-Mall/rpc_gen/kitex_gen/payment"
	"Go-Mall/rpc_gen/kitex_gen/product"
	"context"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"strconv"
)

type CheckoutService struct {
	ctx context.Context
} // NewCheckoutService new CheckoutService
func NewCheckoutService(ctx context.Context) *CheckoutService {
	return &CheckoutService{ctx: ctx}
}

// Run create note info
func (s *CheckoutService) Run(req *checkout.CheckoutReq) (resp *checkout.CheckoutResp, err error) {
	// Finish your business logic.
	// 获取购物车
	cartResult, err := rpc.CartClient.GetCart(s.ctx, &cart.GetCartReq{UserId: req.UserId})
	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(50050001, err.Error())
	}
	if cartResult == nil || cartResult.Items == nil {
		return nil, kerrors.NewGRPCBizStatusError(50050002, "cart is empty")
	}

	var orderItems []*order.OrderItem
	// 计算总价 获取购物车物品
	var total float32
	for _, cartItems := range cartResult.Items {
		productResp, resultErr := rpc.ProductClient.GetProduct(s.ctx, &product.GetProductReq{
			Id: cartItems.ProductId,
		})
		if resultErr != nil {
			return nil, resultErr
		}
		if productResp.Product == nil {
			continue
		}

		p := productResp.Product.Price
		cost := p * float32(cartItems.Quantity)
		total += cost

		orderItem := &order.OrderItem{
			Item: &cart.CartItem{
				ProductId: cartItems.ProductId,
				Quantity:  cartItems.Quantity,
			},
			Cost: cost,
		}
		orderItems = append(orderItems, orderItem)
	}

	// 创建订单
	var orderId string
	zipcode, _ := strconv.ParseInt(req.Address.ZipCode, 10, 32)
	placeOrderResp, err := rpc.OrderClient.PlaceOrder(s.ctx, &order.PlaceOrderReq{
		UserId: req.UserId,
		Address: &order.Address{
			StreetAddress: req.Address.StreetAddress,
			City:          req.Address.City,
			State:         req.Address.State,
			Country:       req.Address.Country,
			ZipCode:       int32(zipcode),
		},
		Email:      req.Email,
		OrderItems: orderItems,
	})
	if err != nil {
		return nil, err
	}
	orderId = placeOrderResp.Order.OrderId

	// 支付请求
	payReq := &payment.ChargeReq{
		UserId:  req.UserId,
		OrderId: orderId,
		Amount:  total,
		CreditCard: &payment.CreditCardInfo{
			CreditCardNumber:          req.CreditCard.CreditCardNumber,
			CreditCardExpirationMonth: req.CreditCard.CreditCardExpirationMonth,
			CreditCardExpirationYear:  req.CreditCard.CreditCardExpirationYear,
			CreditCardCvv:             req.CreditCard.CreditCardCvv,
		},
	}

	// 清空购物车
	_, err = rpc.CartClient.EmptyCart(s.ctx, &cart.EmptyCartReq{UserId: req.UserId})
	if err != nil {
		klog.Error(err.Error())
	}

	//付款
	paymentResult, err := rpc.PaymentClient.Charge(s.ctx, payReq)
	if err != nil {
		return nil, err
	}
	klog.Info(paymentResult)

	//构造响应
	resp = &checkout.CheckoutResp{
		OrderId:       orderId,
		TransactionId: paymentResult.TransactionId,
	}
	return
}
