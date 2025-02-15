package service

import (
	"Go-Mall/app/client/infra/rpc"
	clientutils "Go-Mall/app/client/utils"
	"Go-Mall/rpc_gen/kitex_gen/auth"
	"Go-Mall/rpc_gen/kitex_gen/cart"
	"Go-Mall/rpc_gen/kitex_gen/product"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"strconv"
)

type GetCartService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetCartService(Context context.Context, RequestContext *app.RequestContext) *GetCartService {
	return &GetCartService{RequestContext: RequestContext, Context: Context}
}

func (h *GetCartService) Run() (resp map[string]any, err error) {
	token, err := clientutils.GetTokenFromContext(h.RequestContext)
	if err != nil {
		return nil, err
	}
	// 通过token获取用户信息
	authResp, err := rpc.AuthClient.VerifyTokenByRPC(h.Context, &auth.VerifyTokenReq{Token: token})
	if err != nil {
		return nil, err
	}
	cartResp, err := rpc.CartClient.GetCart(h.Context, &cart.GetCartReq{
		UserId: uint32(authResp.UserId),
	})
	if err != nil {
		return nil, err
	}

	var items []map[string]string
	var total float64
	for _, item := range cartResp.Items {
		productResp, err := rpc.ProductClient.GetProduct(h.Context, &product.GetProductReq{Id: item.ProductId})
		if err != nil {
			continue
		}
		p := productResp.Product
		items = append(items, map[string]string{
			"Name":        p.Name,
			"Description": p.Description,
			"Price":       strconv.FormatFloat(float64(p.Price), 'f', 2, 64),
			"Picture":     p.Picture,
			"Qty":         strconv.Itoa(int(item.Quantity)),
		})
		total += float64(p.Price) * float64(item.Quantity)
	}
	return utils.H{
		"title": "Cart",
		"items": items,
		"total": strconv.FormatFloat(float64(total), 'f', 2, 64),
	}, nil
}
