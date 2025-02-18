package order

import (
	order "Go-Mall/rpc_gen/kitex_gen/order"
	"context"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func PlaceOrder(ctx context.Context, req *order.PlaceOrderReq, callOptions ...callopt.Option) (resp *order.PlaceOrderResp, err error) {
	resp, err = defaultClient.PlaceOrder(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "PlaceOrder call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func ListOrder(ctx context.Context, req *order.ListOrderReq, callOptions ...callopt.Option) (resp *order.ListOrderResp, err error) {
	resp, err = defaultClient.ListOrder(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "ListOrder call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func MarkOrderPaid(ctx context.Context, req *order.MarkOrderPaidReq, callOptions ...callopt.Option) (resp *order.MarkOrderPaidResp, err error) {
	resp, err = defaultClient.MarkOrderPaid(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "MarkOrderPaid call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func GetOrderItems(ctx context.Context, req *order.GetOrderItemsReq, callOptions ...callopt.Option) (resp *order.GetOrderItemsResp, err error) {
	resp, err = defaultClient.GetOrderItems(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "GetOrderItems call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
