package dao

import (
	"context"
	"go-store/model"
	"gorm.io/gorm"
)

type NoticeDao struct {
	*gorm.DB
}

func NewNoticeDao(ctx context.Context) *NoticeDao {
	return &NoticeDao{NewDBClient(ctx)}
}

// NewNoticeDaoByDB 直接使用DB操作的方式
func NewNoticeDaoByDB(db *gorm.DB) *NoticeDao {
	return &NoticeDao{db}
}

// 根据ID获取user
func (d *NoticeDao) GetNoticByID(id int) (notice *model.Notice, err error) {
	err = d.DB.Model(&model.Notice{}).Where("id=?", id).First(&notice).Error
	return
}
