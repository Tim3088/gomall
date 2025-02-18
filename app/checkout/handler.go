package main

import (
	"Go-Mall/app/checkout/biz/service"
	"Go-Mall/rpc_gen/kitex_gen/checkout"
	"context"
)

// CheckoutServiceImpl implements the last service interface defined in the IDL.
type CheckoutServiceImpl struct{}

// Checkout implements the CheckoutServiceImpl interface.
func (s *CheckoutServiceImpl) Checkout(ctx context.Context, req *checkout.CheckoutReq) (resp *checkout.CheckoutResp, err error) {
	resp, err = service.NewCheckoutService(ctx).Run(req)

	return resp, err
}

// CheckoutOrder implements the CheckoutServiceImpl interface.
func (s *CheckoutServiceImpl) CheckoutOrder(ctx context.Context, req *checkout.CheckoutOrderReq) (resp *checkout.CheckoutOrderResp, err error) {
	resp, err = service.NewCheckoutOrderService(ctx).Run(req)

	return resp, err
}
