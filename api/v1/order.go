package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-store/pkg/mytools"
	"go-store/service"
)

func CreateOrder(context *gin.Context) {
	claim, _ := mytools.ParseToken(context.GetHeader("Authorization"))
	createOrderService := service.OrderService{}
	if err := context.ShouldBind(&createOrderService); err == nil {
		res := createOrderService.CreateOrderService(context.Request.Context(), claim.ID)
		context.JSON(200, res)
	} else {
		context.JSON(400, MyError(err))
		mytools.Logger.Infoln("err: ", err)
	}
}

func DeleteOrder(context *gin.Context) {
	claim, _ := mytools.ParseToken(context.GetHeader("Authorization"))
	deleteOrder := service.OrderService{}
	if err := context.ShouldBind(&deleteOrder); err == nil {
		fmt.Println("商品ID", context.Param("id"))
		res := deleteOrder.DeleteOrder(context.Request.Context(), claim.ID, context.Param("id"))
		context.JSON(200, res)
	} else {
		context.JSON(400, MyError(err))
		mytools.Logger.Infoln("err: ", err)
	}
}

func ShowOrder(context *gin.Context) {
	showOrder := service.OrderService{}
	claim, _ := mytools.ParseToken(context.GetHeader("Authorization"))
	if err := context.ShouldBind(&showOrder); err == nil {
		res := showOrder.ShowOrderService(context.Request.Context(), claim.ID)
		context.JSON(200, res)
	} else {
		context.JSON(400, MyError(err))
		mytools.Logger.Infoln("err: ", err)
	}
}

func ShowOrderByID(context *gin.Context) {
	showOrder := service.OrderService{}
	claim, _ := mytools.ParseToken(context.GetHeader("Authorization"))
	if err := context.ShouldBind(&showOrder); err == nil {
		res := showOrder.GetOrderByID(context.Request.Context(), claim.ID, context.Param("id"))
		context.JSON(200, res)
	} else {
		context.JSON(400, MyError(err))
		mytools.Logger.Infoln("err: ", err)
	}
}
