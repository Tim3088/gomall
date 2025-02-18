package service

import (
	"Go-Mall/app/client/infra/rpc"
	clientutils "Go-Mall/app/client/utils"
	"Go-Mall/rpc_gen/kitex_gen/auth"
	rpccheckout "Go-Mall/rpc_gen/kitex_gen/checkout"
	rpcpayment "Go-Mall/rpc_gen/kitex_gen/payment"
	"context"

	checkout "Go-Mall/app/client/hertz_gen/client/checkout"
	"github.com/cloudwego/hertz/pkg/app"
)

type CheckoutOrderService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCheckoutOrderService(Context context.Context, RequestContext *app.RequestContext) *CheckoutOrderService {
	return &CheckoutOrderService{RequestContext: RequestContext, Context: Context}
}

type CheckoutOrderResp struct {
	TransactionId string `json:"transaction_id"`
}

func (h *CheckoutOrderService) Run(req *checkout.CheckoutOrderReq) (resp *CheckoutOrderResp, err error) {
	token, err := clientutils.GetTokenFromContext(h.RequestContext)
	if err != nil {
		return nil, err
	}
	// 通过token获取用户信息
	authResp, err := rpc.AuthClient.VerifyTokenByRPC(h.Context, &auth.VerifyTokenReq{Token: token})
	if err != nil {
		return nil, err
	}

	res, err := rpc.CheckoutClient.CheckoutOrder(h.Context, &rpccheckout.CheckoutOrderReq{
		UserId:  uint32(authResp.UserId),
		OrderId: req.OrderId,
		CreditCard: &rpcpayment.CreditCardInfo{
			CreditCardNumber:          req.CardNum,
			CreditCardCvv:             req.Cvv,
			CreditCardExpirationYear:  req.ExpirationYear,
			CreditCardExpirationMonth: req.ExpirationMonth,
		},
	})
	if err != nil {
		return nil, err
	}

	return &CheckoutOrderResp{
		TransactionId: res.TransactionId,
	}, nil
}
