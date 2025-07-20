package dao

import (
	"context"
	"fmt"
	"go-store/model"
	"gorm.io/gorm"
)

type ShoppingCartDao struct {
	*gorm.DB
}

func (d *ShoppingCartDao) CreateShoppingCart(shoppingCart *model.ShoppingCart) error {
	return d.DB.Model(model.ShoppingCart{}).Create(&shoppingCart).Error
}

func (d *ShoppingCartDao) ListShoppingCartByID(uid int) (shoppingCart []*model.ShoppingCart, err error) {
	err = d.DB.Model(model.ShoppingCart{}).Where("user_id =?", uid).Find(&shoppingCart).Error
	return

}

func (d *ShoppingCartDao) UpdateShoppingCartByID(cid int, ShoppingCart *model.ShoppingCart) error {
	return d.DB.Model(model.ShoppingCart{}).
		Where("id =? ", cid).
		Updates(ShoppingCart).
		Error
}

func (d *ShoppingCartDao) DeleteShoppingCartByID(uid, cid int) error {
	result := d.DB.Model(model.ShoppingCart{}).
		Where("id =? and user_id=?", cid, uid).
		Delete(&model.ShoppingCart{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("no shopping cart record found with cid: %d and uid: %d", cid, uid)
	}
	return nil
}

func (d *ShoppingCartDao) UpdateShoppingNumber(uid, cid, newNum int) error {
	return d.DB.Model(model.ShoppingCart{}).
		Where("id =? and user_id=?", cid, uid).
		Update("the_number_of_units", newNum).
		Error
}

func NewShoppingCartDao(ctx context.Context) *ShoppingCartDao {
	return &ShoppingCartDao{NewDBClient(ctx)}
}
