package dao

import (
	"context"
	"errors"
	"fmt"
	"go-store/model"
	"gorm.io/gorm"
	"log"
)

type OrderDao struct {
	*gorm.DB
}

// CreateOrder 创建订单
func (d *OrderDao) CreateOrder(order *model.Order) error {
	// 直接使用 Create 方法，无需指定 Model
	return d.DB.Create(&order).Error
}

// GetOrderByID 根据订单 ID 获取订单信息
func (d *OrderDao) GetOrderByID(id, uid int) (order *model.Order, err error) {
	// 初始化 order 指针

	err = d.DB.Where("id = ? and user_id=?", id, uid).First(&order).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		order = nil
	}
	return
}

// ListOrderByID 根据用户 ID 获取订单列表
func (d *OrderDao) ListOrderByID(id int) (order []*model.Order, err error) {
	// 使用指针切片接收结果
	err = d.DB.Where("user_id = ?", id).Find(&order).Error
	return
}

// UpdateOrderByID 根据用户 ID 和订单 ID 更新订单信息
func (d *OrderDao) UpdateOrderByID(oid int, uid int, order *model.Order) error {
	// 避免更新零值字段，使用 Select 指定要更新的字段
	return d.DB.Model(&model.Order{}).Where("id = ? and user_id = ?", oid, uid).Select("*").Updates(order).Error
}

// DeleteOrderByID 根据订单 ID 和用户 ID 删除订单
func (d *OrderDao) DeleteOrderByID(aid int, uid int) error {
	result := d.DB.Where("id = ? AND user_id = ?", aid, uid).Delete(&model.Order{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("订单不存在，订单 ID: %d，用户 ID: %d", aid, uid)
	}
	return nil
}

// ListOrderCondition 根据条件分页查询订单列表
func (d *OrderDao) ListOrderCondition(condition map[string]interface{}, page model.BasePage) (order []*model.Order, err error, count int64) {
	// 校验分页参数
	if page.PageNum < 1 {
		page.PageNum = 1
	}
	if page.PageSize < 1 {
		page.PageSize = 10 // 默认每页显示 10 条记录
	}

	// 计算偏移量
	offset := (page.PageNum - 1) * page.PageSize
	if offset < 0 {
		offset = 0
	}

	// 查询总记录数
	query := d.DB.Model(&model.Order{}).Where(condition)
	if err = query.Count(&count).Error; err != nil {
		log.Printf("查询订单总记录数失败: %v", err)
		return nil, err, 0
	}

	// 分页查询订单列表
	if err = query.Offset(offset).Limit(page.PageSize).Find(&order).Error; err != nil {
		log.Printf("分页查询订单列表失败: %v", err)
		return nil, err, count
	}

	return order, nil, count
}

// NewOrderDao 创建一个新的 OrderDao 实例
func NewOrderDao(ctx context.Context) *OrderDao {
	// 假设 NewDBClient 是一个获取数据库连接的函数
	return &OrderDao{NewDBClient(ctx)}
}
