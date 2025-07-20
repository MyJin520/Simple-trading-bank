package serializer

import "go-store/model"

type Carousel struct {
	ID        uint   `json:"id"`
	ImagePath string `json:"image_path"`
	ProductID string `json:"product_id"`
	CreatedAt int64  `json:"created_at"`
}

func BuildCarousel(item *model.Carousel) Carousel {
	return Carousel{
		ID:        item.ID,
		ImagePath: item.ImgPath,
		ProductID: item.ProductID,
		CreatedAt: item.CreatedAt.Unix(),
	}
}

func BuildCarousels(items []model.Carousel) (carousel []Carousel) {
	for _, item := range items {
		carousel = append(carousel, BuildCarousel(&item))
	}
	return carousel
}
