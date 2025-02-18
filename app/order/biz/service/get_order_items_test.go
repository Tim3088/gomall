package service

import (
	order "Go-Mall/rpc_gen/kitex_gen/order"
	"context"
	"testing"
)

func TestGetOrderItems_Run(t *testing.T) {
	ctx := context.Background()
	s := NewGetOrderItemsService(ctx)
	// init req and assert value

	req := &order.GetOrderItemsReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
