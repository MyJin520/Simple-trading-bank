package service

import (
	"context"
	"fmt"
	"go-store/dao"
	"go-store/model"
	"go-store/pkg/e"
	"go-store/pkg/mytools"
	"go-store/serializer"
	"net/http"
	"strconv"
)

type FavoritesService struct {
	ProductID  int `form:"product_id" json:"product_id"`
	BossId     int `json:"boss_id" form:"boss_id"`
	FavoriteId int `json:"favorite_id" form:"favorite_id"`
	model.BasePage
}

func (s *FavoritesService) ShowFavorites(ctx context.Context, uid int) serializer.Response {
	favoriteDao := dao.NewFavoriteDao(ctx)
	code := http.StatusOK
	favorite, err := favoriteDao.ListFavorite(uid)
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
		Data:   serializer.BuildListCarouse(serializer.BuildFavorites(ctx, favorite), int64(len(favorite))),
		Msg:    e.GetMSG(code),
	}
}

func (s *FavoritesService) CreateFavoritesService(ctx context.Context, uid int) interface{} {
	var (
		err1    error
		err2    error
		err3    error
		user    *model.User
		boss    *model.User
		product *model.Product
	)
	favoriteDao := dao.NewFavoriteDao(ctx)
	code := http.StatusOK
	exits, _ := favoriteDao.ExitsOrNot(uid, s.ProductID)
	if exits {
		return serializer.Response{
			Status: http.StatusBadRequest,
			Msg:    e.GetMSG(code),
			Error:  "收藏商品已经存在",
		}
	}
	userDao := dao.NewUserDao(ctx)
	bossDao := dao.NewUserDao(ctx)
	productDao := dao.NewProductDao(ctx)

	user, err1 = userDao.GetUserByID(uid)
	if res := myErr(err1, code); res != nil {
		return res
	}
	boss, err2 = bossDao.GetUserByID(uid)
	if res := myErr(err2, code); res != nil {
		return res
	}
	product, err3 = productDao.GetProductByID(s.ProductID)
	if res := myErr(err3, code); res != nil {
		return res
	}

	favorite := &model.Favorites{
		User:         *user,
		UserID:       uid,
		Product:      *product,
		ProductID:    s.ProductID,
		MerchantName: boss.UserName,
		MerchantID:   s.BossId,
	}
	err := favoriteDao.CreateFavorite(favorite)
	if res := myErr(err, code); res != nil {
		return res
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMSG(code),
	}
}

func (s *FavoritesService) DeleteFavorites(ctx context.Context, id int, fid string) interface{} {

	favoriteDao := dao.NewFavoriteDao(ctx)
	code := http.StatusOK
	fidI, _ := strconv.Atoi(fid)
	fmt.Println(fidI)
	err := favoriteDao.DeleteFavorite(id, fidI)
	if res := myErr(err, code); res != nil {
		return res
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMSG(code),
	}
}

func myErr(err error, code int) interface{} {
	if err != nil {
		mytools.Logger.Infoln("err: ", err)
		code = http.StatusBadRequest
		return serializer.Response{
			Status: code,
			Msg:    e.GetMSG(code),
			Error:  err.Error(),
		}
	}
	return nil
}
