package model

import (
	"context"
	"gorm.io/gorm"
)

type Consignee struct {
	Email         string
	StreetAddress string
	City          string
	State         string
	Country       string
	ZipCode       int32
}

type Order struct {
	gorm.Model
	OrderId      string      `gorm:"type:varchar(100);uniqueIndex"`
	UserId       uint32      `gorm:"type:int(11)"`
	UserCurrency string      `gorm:"type:varchar(10)"`
	Consignee    Consignee   `gorm:"embedded"`
	OrderItems   []OrderItem `gorm:"foreignKey:OrderIdRefer;references:OrderId"`
	Paid         bool        `gorm:"type:boolean;default:false"`
	Firstname    string      `gorm:"type:varchar(100);column:firstname"`
	Lastname     string      `gorm:"type:varchar(100);column:lastname"`
}

func (Order) TableName() string {
	return "order"
}
func ListOrder(ctx context.Context, db *gorm.DB, userId uint32) ([]*Order, error) {
	var orders []*Order
	err := db.WithContext(ctx).Model(&Order{}).Where(&Order{UserId: userId}).Preload("OrderItems").Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func GetOrder(ctx context.Context, db *gorm.DB, orderId string) (*Order, error) {
	var order *Order
	err := db.WithContext(ctx).Model(&Order{}).Where(&Order{OrderId: orderId}).Preload("OrderItems").First(&order).Error
	if err != nil {
		return nil, err
	}
	return order, nil
}
