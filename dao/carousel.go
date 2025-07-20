package dao

import (
	"context"
	"go-store/model"
	"gorm.io/gorm"
)

type CarouselDao struct {
	*gorm.DB
}

func NewCarouselDao(ctx context.Context) *CarouselDao {
	return &CarouselDao{NewDBClient(ctx)}
}

// NewCarouselDaoByDB 直接使用DB操作的方式
func NewCarouselDaoByDB(db *gorm.DB) *CarouselDao {
	return &CarouselDao{db}
}

// GetCarouselDaoByID 根据ID获取user
func (d *CarouselDao) GetCarouselDaoByID(id int) (carousel *model.Carousel, err error) {
	err = d.DB.Model(&model.Carousel{}).Where("id=?", id).First(&carousel).Error
	return
}

func (d *CarouselDao) ListCarousel() (carousel []model.Carousel, err error) {
	err = d.DB.Model(&model.Carousel{}).Find(&carousel).Error
	return
}
