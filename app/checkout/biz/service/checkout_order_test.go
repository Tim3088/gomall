package service

import (
	checkout "Go-Mall/rpc_gen/kitex_gen/checkout"
	"context"
	"testing"
)

func TestCheckoutOrder_Run(t *testing.T) {
	ctx := context.Background()
	s := NewCheckoutOrderService(ctx)
	// init req and assert value

	req := &checkout.CheckoutOrderReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
