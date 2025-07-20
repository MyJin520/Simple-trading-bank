package dao

import (
	"context"
	"go-store/model"
	"gorm.io/gorm"
)

type UserDao struct {
	*gorm.DB
}

// ExistOrNotByUserName 根据userName判断是否存在该用户名
func (d *UserDao) ExistOrNotByUserName(name string) (user *model.User, exist bool, err error) {
	var count int64
	err = d.DB.Model(&model.User{}).Where("user_name=?", name).Find(&user).Count(&count).Error
	if count == 0 {
		return nil, false, err
	}
	return user, true, nil
}

func (d *UserDao) CreateUser(user *model.User) error {
	return d.DB.Model(&model.User{}).Create(&user).Error
}

func (d *UserDao) GetUserByID(id int) (user *model.User, err error) {
	err = d.DB.Model(&model.User{}).Where("id=?", id).First(&user).Error
	return
}

func (d *UserDao) UpdateUserByID(id int, user *model.User) error {
	return d.DB.Model(&model.User{}).Where("id=?", id).Updates(&user).Error
}

func NewUserDao(ctx context.Context) *UserDao {
	return &UserDao{NewDBClient(ctx)}
}

// NewUserDaoByDB 直接使用DB操作的方式
func NewUserDaoByDB(db *gorm.DB) *UserDao {
	return &UserDao{db}
}
