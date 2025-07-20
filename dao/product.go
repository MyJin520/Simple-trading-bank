package dao

import (
	"context"
	"go-store/model"
	"gorm.io/gorm"
)

type ProductDao struct {
	*gorm.DB
}

func NewProductDao(ctx context.Context) *ProductDao {
	return &ProductDao{NewDBClient(ctx)}
}

func NewProductDaoByDB(db *gorm.DB) *ProductDao {
	return &ProductDao{db}
}

func (pd *ProductDao) CreateProduct(product *model.Product) error {
	// 直接传入 product 指针
	return pd.DB.Model(&model.Product{}).Create(&product).Error
}

func (pd *ProductDao) CountProductByCondition(condition map[string]interface{}) (total int64, err error) {
	err = pd.DB.Model(&model.Product{}).Where(condition).Count(&total).Error
	return
}

func (pd *ProductDao) ListProductByCondition(condition map[string]interface{}, page model.BasePage) (product []*model.Product, err error) {
	err = pd.DB.Where(condition).Offset(page.PageSize * (page.PageNum - 1)).Limit(page.PageSize).Find(&product).Error
	return
}

func (pd *ProductDao) FindProduct(info string, page model.BasePage) (product []*model.Product, count int64, err error) {
	err = pd.DB.Model(&model.Product{}).Where("info LIKE?", "%"+info+"%").
		Count(&count).Error
	if err != nil {
		return
	}
	err = pd.DB.Where("info LIKE ?", "%"+info+"%").
		Offset(page.PageSize * (page.PageNum - 1)).
		Limit(page.PageSize).
		Find(&product).Error
	return
}

func (pd *ProductDao) GetProductByID(id int) (product *model.Product, err error) {
	err = pd.DB.Model(&model.Product{}).Where("id = ?", id).First(&product).Error
	return

}

func (pd *ProductDao) UpdateProduct(pid int, product *model.Product) error {
	return pd.DB.Model(&model.Product{}).Where("id =?", pid).Updates(product).Error
}
