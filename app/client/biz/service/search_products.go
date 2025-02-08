package service

import (
	"Go-Mall/app/client/infra/rpc"
	rpcproduct "Go-Mall/rpc_gen/kitex_gen/product"
	"context"

	product "Go-Mall/app/client/hertz_gen/client/product"
	"github.com/cloudwego/hertz/pkg/app"
)

type SearchProductsService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

type SearchProductResp struct {
	Products []*rpcproduct.Product `json:"products"`
}

func NewSearchProductsService(Context context.Context, RequestContext *app.RequestContext) *SearchProductsService {
	return &SearchProductsService{RequestContext: RequestContext, Context: Context}
}

func (h *SearchProductsService) Run(req *product.SearchProductsReq) (resp *SearchProductResp, err error) {
	res, err := rpc.ProductClient.SearchProducts(h.Context, &rpcproduct.SearchProductsReq{Query: req.GetQ()})
	if err != nil {
		return nil, err
	}
	return &SearchProductResp{Products: res.Results}, nil
}
