package service

import (
	"Go-Mall/app/client/hertz_gen/client/product"
	"Go-Mall/app/client/infra/rpc"
	rpcproduct "Go-Mall/rpc_gen/kitex_gen/product"
	"context"

	"github.com/cloudwego/hertz/pkg/app"
)

type ListProductsService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewListProductsService(Context context.Context, RequestContext *app.RequestContext) *ListProductsService {
	return &ListProductsService{RequestContext: RequestContext, Context: Context}
}

type ListProductResp struct {
	Products []*rpcproduct.Product `json:"products"`
}

func (h *ListProductsService) Run(req *product.ListProductsReq) (resp *ListProductResp, err error) {
	res, err := rpc.ProductClient.ListProducts(h.Context, &rpcproduct.ListProductsReq{Page: req.GetPage(), PageSize: req.GetPageSize(), CategoryName: req.GetCategoryName()})
	if err != nil {
		return nil, err
	}
	return &ListProductResp{Products: res.Products}, nil
}
