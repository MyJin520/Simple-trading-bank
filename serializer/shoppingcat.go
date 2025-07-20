package serializer

import (
	"context"
	"go-store/conf"
	"go-store/dao"
	"go-store/model"
)

type ShoppingCart struct {
	ID               int    `json:"id" form:"id"`
	UserID           int    `json:"user_id" form:"user_id"`
	ProductID        int    `json:"product_id" form:"product_id"`
	CreatedAt        int64  `json:"created_at" form:"created_at"`
	TheNumberOfUnits int    `json:"the_number_of_units" form:"the_number_of_units"` // 商品数量
	MaximumPurchase  uint   `json:"maximum_purchase" form:"maximum_purchase"`       // 最大购买数
	PaymentChecks    bool   `json:"payment_checks" form:"payment_checks"`           // 支付检查
	ImagePath        string `json:"image_path" form:"image_path"`                   // 商品图片
	ProductName      string `json:"product_name" form:"product_name"`               // 商品名称
	DiscountPrice    string `json:"discount_price" form:"discount_price"`           // 折扣
	BossID           int    `json:"boss_id" form:"boss_id"`                         // 商家ID
	BossName         string `json:"boss_name" form:"boss_name"`                     // 商家名称
}

func BuildShoppingCart(cart *model.ShoppingCart, product *model.Product) ShoppingCart {
	return ShoppingCart{
		ID:               int(cart.ID),
		UserID:           cart.UserID,
		ProductID:        cart.ProductID,
		CreatedAt:        cart.CreatedAt.Unix(),
		TheNumberOfUnits: cart.TheNumberOfUnits,
		MaximumPurchase:  cart.MaximumPurchase,
		PaymentChecks:    cart.PaymentChecks,
		ProductName:      product.Name,
		ImagePath:        conf.Host + conf.HttpPort + conf.ProductPath + product.ImgPath,
		DiscountPrice:    product.DiscountPerice,
		BossID:           product.MerchantID,
		BossName:         product.MerchantName,
	}
}

func BuildShoppingCarts(ctx context.Context, items []*model.ShoppingCart) (cart []ShoppingCart) {
	productDao := dao.NewProductDao(ctx)
	bossDao := dao.NewUserDao(ctx)
	for _, item := range items {
		product, err := productDao.GetProductByID(item.ProductID)
		if err != nil {
			continue
		}
		_, err = bossDao.GetUserByID(item.MerchantID)
		if err != nil {
			continue
		}
		cart = append(cart, BuildShoppingCart(item, product))
	}
	return cart
}
