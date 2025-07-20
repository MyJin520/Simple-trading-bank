package model

import "gorm.io/gorm"

// Favorites 收藏夹
type Favorites struct {
	gorm.Model
	User         User    // 根据用户ID,此处未使用外键
	UserID       int     `gorm:"not null"`
	Product      Product // 根据商品ID,此处未使用外键
	ProductID    int     `gorm:"not null"`
	MerchantName string  // 商家名称【boss】
	MerchantID   int     // 商家ID

}
