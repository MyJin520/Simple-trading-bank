package model

import (
	"go-store/cache"
	"gorm.io/gorm"
	"strconv"
)

type Product struct {
	gorm.Model
	Name           string
	CategoryID     int    // 类别
	Title          string // 商品标贴
	Info           string // 商品介绍
	ImgPath        string
	Price          string // 商品价格
	DiscountPerice string // 打折后的价格
	OnSale         bool   `gorm:"default:false"` // 是否在售
	Number         int    // 再售数量
	MerchantID     int    // 商家ID
	MerchantName   string // 商家名称
	MerchantAvatar string
}

func (p *Product) ViewNumber() int64 {
	countStr, _ := cache.RedisClient.Get(cache.ProductViewKey(p.ID)).Result()
	count, _ := strconv.ParseInt(countStr, 10, 64)
	return count
}

func (p *Product) AddView() {
	// 增加商品点击数
	cache.RedisClient.Incr(cache.ProductViewKey(p.ID))
	cache.RedisClient.ZIncrBy(cache.RankKey, 1, strconv.Itoa(int(p.ID)))
}
