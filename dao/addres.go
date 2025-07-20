package dao

import (
	"context"
	"go-store/model"
	"gorm.io/gorm"
)

type AddressDao struct {
	*gorm.DB
}

func (d *AddressDao) CreateAddress(address *model.Address) error {
	return d.DB.Model(model.Address{}).Create(&address).Error
}

func (d *AddressDao) GetAddressByID(id int) (address *model.Address, err error) {
	err = d.DB.Model(model.Address{}).Where("id = ?", id).First(&address).Error
	return
}

func (d *AddressDao) ListAddressByID(id int) (address []*model.Address, err error) {
	err = d.DB.Model(model.Address{}).Where("user_id =?", id).Find(&address).Error
	return

}

func (d *AddressDao) UpdateAddressByID(uid int, aid int, address *model.Address) error {
	return d.DB.Model(model.Address{}).Where("id =? and user_id=?", aid, uid).Updates(address).Error
}

func (d *AddressDao) DeleteAddressByID(aid int, uid int) error {
	return d.DB.Model(model.Address{}).Where("id =? and user_id=?", aid, uid).Delete(&model.Address{}).Error
}

func NewAddressDao(ctx context.Context) *AddressDao {
	return &AddressDao{NewDBClient(ctx)}
}
