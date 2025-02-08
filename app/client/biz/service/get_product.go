package service

import (
	"Go-Mall/app/client/infra/rpc"
	rpcproduct "Go-Mall/rpc_gen/kitex_gen/product"
	"context"

	product "Go-Mall/app/client/hertz_gen/client/product"
	"github.com/cloudwego/hertz/pkg/app"
)

type GetProductService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

type GetProductResp struct {
	Product *rpcproduct.Product `json:"product"`
}

func NewGetProductService(Context context.Context, RequestContext *app.RequestContext) *GetProductService {
	return &GetProductService{RequestContext: RequestContext, Context: Context}
}

func (h *GetProductService) Run(req *product.ProductReq) (resp *GetProductResp, err error) {
	res, err := rpc.ProductClient.GetProduct(h.Context, &rpcproduct.GetProductReq{Id: req.GetId()})
	if err != nil {
		return nil, err
	}
	return &GetProductResp{Product: res.Product}, nil
}
