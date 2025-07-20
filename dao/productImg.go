package dao

import (
	"context"
	"go-store/model"
	"go-store/pkg/mytools"
	"gorm.io/gorm"
)

type ProductImagDao struct {
	*gorm.DB
}

func NewProductImageDao(ctx context.Context) *ProductImagDao {
	return &ProductImagDao{NewDBClient(ctx)}
}

func NewProductImageDaoByID(db *gorm.DB) *ProductImagDao {
	return &ProductImagDao{db}
}

func (dao *ProductImagDao) CreateProductImage(productImage *model.ProductImg) error {
	err := dao.DB.Model(&model.ProductImg{}).Create(&productImage).Error
	if err != nil {
		mytools.Logger.Infof("插入商品图片信息失败，错误信息: %v，商品图片信息: %+v", err, productImage)
	}
	return err
}
