package model

import (
	"context"
	"gorm.io/gorm"
)

type Category struct {
	Base
	Name        string `json:"name"`
	Description string `json:"description"`

	Products []Product `json:"products" gorm:"many2many:product_category"`
}

func (c Category) TableName() string {
	return "category"
}

type CategoryQuery struct {
	ctx context.Context
	db  *gorm.DB
}

func (c CategoryQuery) GetProductsByCategoryName(name string, page int, pageSize int) (products []Product, err error) {
	// 计算分页的偏移量
	offset := (page - 1) * pageSize

	// 直接查询指定分类下的产品，并应用分页
	err = c.db.WithContext(c.ctx).Model(&Product{}).
		Joins("JOIN product_category ON product.id = product_category.product_id").
		Joins("JOIN category ON product_category.category_id = category.id").
		Where("category.name = ?", name).
		Offset(offset).Limit(pageSize).Find(&products).Error
	return
}

func NewCategoryQuery(ctx context.Context, db *gorm.DB) *CategoryQuery {
	return &CategoryQuery{
		ctx: ctx,
		db:  db,
	}
}
