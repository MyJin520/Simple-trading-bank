package service

import (
	"context"
	"go-store/dao"
	"go-store/pkg/e"
	"go-store/pkg/mytools"
	"go-store/serializer"
	"net/http"
)

type CategoryService struct {
	ID int `form:"id"`
}

func (s *CategoryService) List(ctx context.Context) serializer.Response {
	CategoryDao := dao.NewCategoryDao(ctx)
	code := http.StatusOK
	category, err := CategoryDao.ListCategory()
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
		Data:   serializer.BuildListCarouse(serializer.BuildCategorys(category), int64(len(category))),
		Msg:    e.GetMSG(code),
	}
}
