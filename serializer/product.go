package serializer

import (
	"go-store/conf"
	"go-store/model"
)

type SerializeProduct struct {
	ID            uint   `json:"id"`
	Name          string `json:"name"`
	CategoryID    int    `json:"category_id"`
	Title         string `json:"title"`
	Info          string `json:"info"`
	ImgPath       string `json:"img_path"`
	Price         string `json:"price"`
	DiscountPrice string `json:"discount_price"`
	OnSale        bool   `json:"on_sale"`
	Num           int    `json:"num"`
	ViewNumber    int64  `json:"view_number"` // 商品浏览数量
	BossID        int    `json:"boss_id"`     // 商家ID
	BossName      string `json:"boss_name"`   // 商家ID
	BossAvatar    string `json:"boss_avatar"` // 商家头像
	CreatedAt     int64  `json:"created_at"`  // 创建时间
}

func BuildProduct(product *model.Product) SerializeProduct {
	return SerializeProduct{
		ID:            product.ID,
		Name:          product.Name,
		CategoryID:    product.CategoryID, // 类别ID
		Title:         product.Title,
		Info:          product.Info,
		ImgPath:       conf.Host + conf.HttpPort + conf.ProductPath + product.ImgPath,
		Price:         product.Price,
		DiscountPrice: product.DiscountPerice,
		OnSale:        product.OnSale,
		Num:           product.Number,
		ViewNumber:    product.ViewNumber(),
		BossID:        product.MerchantID,
		BossName:      product.MerchantName,
		BossAvatar:    product.MerchantAvatar,
		CreatedAt:     product.CreatedAt.Unix(),
	}
}

func BuildProducts(items []*model.Product) (products []SerializeProduct) {
	for _, item := range items {
		product := BuildProduct(item)
		products = append(products, product)
	}
	return products
}
