package serializer

import "go-store/model"

type Category struct {
	ID           uint   `json:"id"`
	CategoryName string `json:"category_name"`
	CreatedAt    int64  `json:"created_at"`
}

func BuildCategory(item *model.ProductCategory) Category {
	return Category{
		ID:           item.ID,
		CategoryName: item.CategoryName,
		CreatedAt:    item.CreatedAt.Unix(),
	}
}

func BuildCategorys(items []*model.ProductCategory) (Category []Category) {
	for _, item := range items {
		Category = append(Category, BuildCategory(item))
	}
	return Category
}
