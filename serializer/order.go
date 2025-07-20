package serializer

import (
	"context"
	"fmt"
	"go-store/conf"
	"go-store/dao"
	"go-store/model"
)

type OrderSerializer struct {
	ID            uint    `json:"id" form:"id"`                         // 订单ID
	OrderNum      int     `json:"order_num" form:"order_num"`           // 订单编号
	CreateAt      int64   `json:"create_at" form:"create_at"`           // 订单创建时间
	UpdateAt      int64   `json:"update_at" form:"update_at"`           // 订单更新时间
	UserName      string  `json:"user_name" form:"user_name"`           // 用户名
	ProductName   string  `json:"product_name" form:"product_name"`     // 商品名称
	Address       string  `json:"address" form:"address"`               // 收货地址
	Phone         string  `json:"phone" form:"phone"`                   // 联系电话
	BossName      string  `json:"boss_name" form:"boss_name"`           // 商家名称
	Price         float64 `json:"price" form:"price"`                   // 订单价格
	Type          uint    `json:"type" form:"type"`                     // 订单是否支付 0 未支付 1 已支付
	ImagePath     string  `json:"image_path" form:"image_path"`         // 商品图片路径
	DiscountPrice string  `json:"discount_price" form:"discount_price"` // 优惠后价格
}

func BuildOrder(order *model.Order, product *model.Product, address *model.Address, user, boss *model.User) OrderSerializer {
	fmt.Println("DiscountPrice", product.DiscountPerice, "Price", product.Price)
	return OrderSerializer{
		ID:            order.ID,
		OrderNum:      order.OrderNum,
		CreateAt:      order.CreatedAt.Unix(),
		UpdateAt:      order.UpdatedAt.Unix(),
		UserName:      user.UserName,
		ProductName:   product.Name,
		Address:       address.Address,
		Phone:         address.Phone,
		BossName:      boss.UserName,
		Type:          order.Type,
		ImagePath:     conf.Host + conf.HttpPort + conf.ProductPath + product.ImgPath,
		DiscountPrice: product.DiscountPerice,
		Price:         order.Money,
	}
}

func BuildOrders(ctx context.Context, items []*model.Order) (order []OrderSerializer) {
	productDao := dao.NewProductDao(ctx)
	bossDao := dao.NewUserDao(ctx)
	address := dao.NewAddressDao(ctx)
	user := dao.NewUserDao(ctx)

	for _, item := range items {
		product, err1 := productDao.GetProductByID(item.ProductID)
		if err1 != nil {
			continue
		}
		boss, err2 := bossDao.GetUserByID(item.MerchantID)
		if err2 != nil {
			continue
		}
		address, err3 := address.GetAddressByID(item.AddressID)
		if err3 != nil {
			continue
		}
		user, err4 := user.GetUserByID(item.UserID)
		if err4 != nil {
			continue
		}
		order = append(order, BuildOrder(item, product, address, user, boss))
	}
	return order
}
