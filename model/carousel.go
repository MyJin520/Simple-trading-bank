package model

import (
	"gorm.io/gorm"
	"time"
)

// Carousel 轮播图
type Carousel struct {
	gorm.Model
	ImgPath   string // 图片路径
	ImgID     int    `gorm:"not null"`
	ProductID string
	CreatedAt time.Time
}
