package model

import "gorm.io/gorm"

type ProductImg struct {
	gorm.Model
	ProductID int `gorm:"not null"`
	ImgPath   string
}
