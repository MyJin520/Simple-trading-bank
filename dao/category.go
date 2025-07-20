package dao

import (
	"context"
	"go-store/model"
	"gorm.io/gorm"
)

type CategoryDao struct {
	*gorm.DB
}

func NewCategoryDao(ctx context.Context) *CategoryDao {
	return &CategoryDao{NewDBClient(ctx)}
}

// NewCategoryDaoByDB 直接使用DB操作的方式
func NewCategoryDaoByDB(db *gorm.DB) *CategoryDao {
	return &CategoryDao{db}
}

// GetCategoryDaoByID 根据ID获取user
func (d *CategoryDao) GetCategoryDaoByID(id int) (Category *model.ProductCategory, err error) {
	err = d.DB.Model(&model.ProductCategory{}).Where("id=?", id).First(&Category).Error
	return
}

func (d *CategoryDao) ListCategory() (Category []*model.ProductCategory, err error) {
	err = d.DB.Model(&model.ProductCategory{}).Find(&Category).Error
	return
}
