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

type CheckoutService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCheckoutService(Context context.Context, RequestContext *app.RequestContext) *CheckoutService {
	return &CheckoutService{RequestContext: RequestContext, Context: Context}
}

type CheckoutResp struct {
	OrderId       string `json:"order_id"`
	TransactionId string `json:"transaction_id"`
}

func (h *CheckoutService) Run(req *checkout.CheckoutReq) (resp *CheckoutResp, err error) {
	token, err := clientutils.GetTokenFromContext(h.RequestContext)
	if err != nil {
		return nil, err
	}
	// 通过token获取用户信息
	authResp, err := rpc.AuthClient.VerifyTokenByRPC(h.Context, &auth.VerifyTokenReq{Token: token})
	if err != nil {
		return nil, err
	}

	res, err := rpc.CheckoutClient.Checkout(h.Context, &rpccheckout.CheckoutReq{
		UserId:    uint32(authResp.UserId),
		Firstname: req.Firstname,
		Lastname:  req.Lastname,
		Email:     req.Email,
		Address: &rpccheckout.Address{
			State:         req.Province,
			City:          req.City,
			StreetAddress: req.Street,
			ZipCode:       req.Zipcode,
			Country:       req.Country,
		},
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
	return &CheckoutResp{
		OrderId:       res.OrderId,
		TransactionId: res.TransactionId,
	}, nil
}
