package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-store/pkg/mytools"
	"go-store/service"
)

func CreateShoppingCart(context *gin.Context) {
	claim, _ := mytools.ParseToken(context.GetHeader("Authorization"))
	createShoppingCartService := service.ShoppingCartService{}
	if err := context.ShouldBind(&createShoppingCartService); err == nil {
		res := createShoppingCartService.CreateShoppingCartService(context.Request.Context(), claim.ID)
		context.JSON(200, res)
	} else {
		context.JSON(400, MyError(err))
		mytools.Logger.Infoln("err: ", err)
	}
}

func DeleteShoppingCart(context *gin.Context) {
	claim, _ := mytools.ParseToken(context.GetHeader("Authorization"))
	deleteShoppingCart := service.ShoppingCartService{}
	if err := context.ShouldBind(&deleteShoppingCart); err == nil {
		fmt.Println("商品ID", context.Param("id"))
		res := deleteShoppingCart.DeleteShoppingCart(context.Request.Context(), claim.ID, context.Param("id"))
		context.JSON(200, res)
	} else {
		context.JSON(400, MyError(err))
		mytools.Logger.Infoln("err: ", err)
	}
}

func ShowShoppingCart(context *gin.Context) {
	showShoppingCart := service.ShoppingCartService{}
	claim, _ := mytools.ParseToken(context.GetHeader("Authorization"))
	if err := context.ShouldBind(&showShoppingCart); err == nil {
		res := showShoppingCart.ShowShoppingCart(context.Request.Context(), claim.ID)
		context.JSON(200, res)
	} else {
		context.JSON(400, MyError(err))
		mytools.Logger.Infoln("err: ", err)
	}
}

func UpdateShoppingCart(context *gin.Context) {
	showShoppingCart := service.ShoppingCartService{}
	claim, _ := mytools.ParseToken(context.GetHeader("Authorization"))
	if err := context.ShouldBind(&showShoppingCart); err == nil {
		res := showShoppingCart.UpdateShoppingCart(context.Request.Context(), claim.ID, context.Param("id"))
		context.JSON(200, res)
	} else {
		context.JSON(400, MyError(err))
		mytools.Logger.Infoln("err: ", err)
	}
}
