package service

import (
	"context"
	"go-store/dao"
	"go-store/model"
	"go-store/pkg/e"
	"go-store/serializer"
	"net/http"
	"strconv"
)

type ShoppingCartService struct {
	ID         int `json:"id" form:"id"`
	BossId     int `json:"boss_id" form:"boss_id"`
	ProductID  int `json:"product_id" form:"product_id"`
	ProductNum int `json:"product_num" form:"product_num"`
}

func (s *ShoppingCartService) CreateShoppingCartService(ctx context.Context, uid int) serializer.Response {
	// 提前初始化返回的响应对象
	resp := serializer.Response{
		Status: http.StatusOK,
		Msg:    e.GetMSG(http.StatusOK),
	}

	// 检查商品是否存在
	productDao := dao.NewProductDao(ctx)
	product, err := productDao.GetProductByID(s.ProductID)
	if err != nil {
		resp.Status = http.StatusBadRequest
		resp.Msg = e.GetMSG(http.StatusBadRequest)
		resp.Error = err.Error()
		return resp
	}

	// 检查商家是否存在
	userDao := dao.NewUserDao(ctx)
	_, err = userDao.GetUserByID(product.MerchantID)
	if err != nil {
		resp.Status = http.StatusBadRequest
		resp.Msg = e.GetMSG(http.StatusBadRequest)
		resp.Error = "商家不存在"
		return resp
	}

	// 创建购物车记录
	shoppingCartDao := dao.NewShoppingCartDao(ctx)
	shoppingCart := &model.ShoppingCart{
		UserID:     uid,
		ProductID:  s.ProductID,
		MerchantID: product.MerchantID,
	}
	err = shoppingCartDao.CreateShoppingCart(shoppingCart)
	if err != nil {
		resp.Status = http.StatusBadRequest
		resp.Msg = e.GetMSG(http.StatusBadRequest)
		resp.Error = err.Error()
		return resp
	}

	// 设置成功响应数据
	resp.Data = serializer.BuildShoppingCart(shoppingCart, product)
	return resp
}

func (s *ShoppingCartService) ShowShoppingCart(ctx context.Context, uid int) serializer.Response {
	code := http.StatusOK
	ShoppingCartDao := dao.NewShoppingCartDao(ctx)
	ShoppingCartList, err := ShoppingCartDao.ListShoppingCartByID(uid)
	if err != nil {
		code = http.StatusBadRequest
		return serializer.Response{
			Status: code,
			Msg:    e.GetMSG(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildShoppingCarts(ctx, ShoppingCartList),
		Msg:    e.GetMSG(code),
	}
}

func (s *ShoppingCartService) UpdateShoppingCart(ctx context.Context, uid int, cID string) serializer.Response {
	shoppingCartDao := dao.NewShoppingCartDao(ctx)
	cartID, _ := strconv.Atoi(cID)
	err := shoppingCartDao.UpdateShoppingNumber(uid, cartID, s.ProductNum)
	if err != nil {
		return serializer.Response{
			Status: http.StatusBadRequest,
			Msg:    e.GetMSG(http.StatusBadRequest),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: http.StatusOK,
		Msg:    e.GetMSG(http.StatusOK),
		Data:   interface{}("更新后购买数量" + strconv.Itoa(s.ProductNum)),
	}
}

func (s *ShoppingCartService) DeleteShoppingCart(ctx context.Context, uid int, cid string) interface{} {
	ShoppingCartDao := dao.NewShoppingCartDao(ctx)
	cID, _ := strconv.Atoi(cid)
	err := ShoppingCartDao.DeleteShoppingCartByID(uid, cID)
	if err != nil {
		return serializer.Response{
			Status: http.StatusBadRequest,
			Msg:    e.GetMSG(http.StatusBadRequest),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: http.StatusOK,
		Msg:    e.GetMSG(http.StatusOK),
	}
}
