package service

import (
	"Go-Mall/app/product/biz/dal/mysql"
	"Go-Mall/app/product/biz/model"
	product "Go-Mall/rpc_gen/kitex_gen/product"
	"context"
)

type ListProductsService struct {
	ctx context.Context
} // NewListProductsService new ListProductsService
func NewListProductsService(ctx context.Context) *ListProductsService {
	return &ListProductsService{ctx: ctx}
}

// Run create note info
func (s *ListProductsService) Run(req *product.ListProductsReq) (resp *product.ListProductsResp, err error) {
	// Finish your business logic.
	categoryQuery := model.NewCategoryQuery(s.ctx, mysql.DB)

	c, err := categoryQuery.GetProductsByCategoryName(
		req.CategoryName,
		int(req.Page),
		int(req.PageSize),
	)

	resp = &product.ListProductsResp{}
	for _, v := range c {
		resp.Products = append(resp.Products, &product.Product{
			Id:          uint32(v.ID),
			Name:        v.Name,
			Description: v.Description,
			Picture:     v.Picture,
			Price:       v.Price,
		})
	}
	return resp, nil
}
