package service

import (
	"context"
	"go-store/dao"
	"go-store/pkg/e"
	"go-store/pkg/mytools"
	"go-store/serializer"
	"net/http"
)

type CarouselService struct {
	ID int `form:"id"`
}

func (s *CarouselService) List(ctx context.Context) interface{} {
	carouselDao := dao.NewCarouselDao(ctx)
	code := http.StatusOK
	carousel, err := carouselDao.ListCarousel()
	if err != nil {
		mytools.Logger.Infoln("err: ", err)
		code = http.StatusBadRequest
		return serializer.Response{
			Status: code,
			Msg:    e.GetMSG(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildListCarouse(serializer.BuildCarousels(carousel), int64(len(carousel))),
		Msg:    e.GetMSG(code),
	}
}
