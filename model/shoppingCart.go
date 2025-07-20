package model

import "gorm.io/gorm"

// ShoppingCart 购物车
type ShoppingCart struct {
	gorm.Model
	UserID           int  `gorm:"not null"`
	ProductID        int  `gorm:"not null"` // 商品ID
	MerchantID       int  `gorm:"not null"` // 商家ID
	TheNumberOfUnits int  `gorm:"not null"` // 商品数量
	MaximumPurchase  uint `gorm:"not null"` // 最大购买数
	PaymentChecks    bool `gorm:"not null"` // 支付检查

}
