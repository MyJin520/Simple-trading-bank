package serializer

import (
	"context"
	"go-store/conf"
	"go-store/dao"
	"go-store/model"
)

type Favorite struct {
	UserID        int    `json:"user_id"`
	ProductID     int    `json:"product_id"`
	CreatedAt     int64  `json:"created_at"`
	Name          string `json:"name"`
	CategoryID    int    `json:"category_id"`
	Title         string `json:"title"`
	Info          string `json:"info"`
	ImagePath     string `json:"image_path"`
	Price         string `json:"price"`
	DiscountPrice string `json:"discount_price"`
	BossID        int    `json:"boss_id"`
	Num           int    `json:"num"`
	OnSale        bool   `json:"on_sale"`
}

func BuildFavorite(favorite *model.Favorites, product *model.Product, boss *model.User) Favorite {
	return Favorite{
		UserID:        favorite.UserID,
		ProductID:     favorite.ProductID,
		CreatedAt:     favorite.CreatedAt.Unix(),
		Name:          product.Name,
		CategoryID:    product.CategoryID,
		Title:         product.Title,
		Info:          product.Info,
		ImagePath:     conf.Host + conf.HttpPort + conf.ProductPath + product.ImgPath,
		Price:         product.Price,
		DiscountPrice: product.DiscountPerice,
		BossID:        int(boss.ID),
		Num:           product.Number,
		OnSale:        product.OnSale,
	}
}

func BuildFavorites(ctx context.Context, items []*model.Favorites) (Favorite []Favorite) {
	productDao := dao.NewProductDao(ctx)
	bossDao := dao.NewUserDao(ctx)
	for _, item := range items {
		product, err := productDao.GetProductByID(item.ProductID)
		if err != nil {
			continue
		}
		boss, err := bossDao.GetUserByID(item.UserID)
		if err != nil {
			continue
		}
		Favorite = append(Favorite, BuildFavorite(item, product, boss))
	}
	return Favorite
}
