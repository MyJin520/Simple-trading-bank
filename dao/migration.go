package dao

import (
	"fmt"
	"go-store/model"
)

func Migration() {
	err := _db.Set("gorm:table_options", "charset=utf8mb4").
		// 自动迁移数据库
		AutoMigrate(
			&model.Address{},
			&model.Admin{},
			&model.BasePage{},
			&model.Carousel{},
			&model.Favorites{},
			&model.Notice{},
			&model.Order{},
			&model.Product{},
			&model.ProductCategory{},
			&model.ProductImg{},
			&model.ShoppingCart{},
			&model.User{},
		)
	if err != nil {
		fmt.Println("err：", err)
		return
	}
	fmt.Println("数据迁移成功！！！")
}
