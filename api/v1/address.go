package v1

import (
	"github.com/gin-gonic/gin"
	"go-store/pkg/mytools"
	"go-store/service"
)

func CreateAddress(context *gin.Context) {
	claim, _ := mytools.ParseToken(context.GetHeader("Authorization"))
	createAddressService := service.AddressService{}
	if err := context.ShouldBind(&createAddressService); err == nil {
		res := createAddressService.CreateAddressService(context.Request.Context(), claim.ID)
		context.JSON(200, res)
	} else {
		context.JSON(400, MyError(err))
		mytools.Logger.Infoln("err: ", err)
	}
}

func DeleteAddress(context *gin.Context) {
	claim, _ := mytools.ParseToken(context.GetHeader("Authorization"))
	deleteAddress := service.AddressService{}
	if err := context.ShouldBind(&deleteAddress); err == nil {
		res := deleteAddress.DeleteAddress(context.Request.Context(), claim.ID, context.Param("id"))
		context.JSON(200, res)
	} else {
		context.JSON(400, MyError(err))
		mytools.Logger.Infoln("err: ", err)
	}
}

func GetByIDAddress(context *gin.Context) {
	showAddress := service.AddressService{}
	if err := context.ShouldBind(&showAddress); err == nil {
		res := showAddress.GetAddressByID(context.Request.Context(), context.Param("id"))
		context.JSON(200, res)
	} else {
		context.JSON(400, MyError(err))
		mytools.Logger.Infoln("err: ", err)
	}
}

func ShowAddress(context *gin.Context) {
	showAddress := service.AddressService{}
	claim, _ := mytools.ParseToken(context.GetHeader("Authorization"))
	if err := context.ShouldBind(&showAddress); err == nil {
		res := showAddress.ShowAddress(context.Request.Context(), claim.ID)
		context.JSON(200, res)
	} else {
		context.JSON(400, MyError(err))
		mytools.Logger.Infoln("err: ", err)
	}
}

func UpdateAddress(context *gin.Context) {
	showAddress := service.AddressService{}
	claim, _ := mytools.ParseToken(context.GetHeader("Authorization"))
	if err := context.ShouldBind(&showAddress); err == nil {
		res := showAddress.UpdateAddress(context.Request.Context(), claim.ID, context.Param("id"))
		context.JSON(200, res)
	} else {
		context.JSON(400, MyError(err))
		mytools.Logger.Infoln("err: ", err)
	}
}
