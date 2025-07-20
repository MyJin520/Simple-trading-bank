package model

import "gorm.io/gorm"

// ProductCategory 商品分类
type ProductCategory struct {
	gorm.Model
	CategoryName string // 类别名称
}
