package model

import "gorm.io/gorm"

// Order 订单
type Order struct {
	gorm.Model
	UserID     int `gorm:"not null"`
	ProductID  int `gorm:"not null"`
	MerchantID int `gorm:"not null"`
	AddressID  int `gorm:"not null"`
	Num        int
	OrderNum   int
	Type       uint // 0未支付，1已支付
	Money      float64
}
