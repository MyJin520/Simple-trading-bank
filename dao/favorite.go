package dao

import (
	"context"
	"go-store/model"
	"gorm.io/gorm"
)

type FavoriteDao struct {
	*gorm.DB
}

func (d *FavoriteDao) ListFavorite(uid int) (resp []*model.Favorites, err error) {
	d.DB.Model(&model.Favorites{}).Where("user_id=?", uid).Find(&resp)
	return
}

func (d *FavoriteDao) ExitsOrNot(uid int, favoriteId int) (exits bool, err error) {
	var count int64
	err = d.DB.Model(&model.Favorites{}).Where("user_id=? and product_id=?", uid, favoriteId).Count(&count).Error
	// 不存在返回false，存在返回true
	if err != nil || count == 0 {
		return false, err
	}
	return true, nil
}

func (d *FavoriteDao) CreateFavorite(favorite *model.Favorites) error {
	return d.DB.Model(&model.Favorites{}).Create(&favorite).Error

}

func (d *FavoriteDao) DeleteFavorite(id int, fid int) error {
	return d.DB.Model(&model.Favorites{}).Where(" product_id=? and user_id=?", fid, id).Delete(&model.Favorites{}).Error

}

func NewFavoriteDao(ctx context.Context) *FavoriteDao {
	return &FavoriteDao{NewDBClient(ctx)}
}

// NewFavoriteDaoByDB 直接使用DB操作的方式
func NewFavoriteDaoByDB(db *gorm.DB) *FavoriteDao {
	return &FavoriteDao{db}
}
