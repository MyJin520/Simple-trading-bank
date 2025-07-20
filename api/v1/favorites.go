package v1

import (
	"github.com/gin-gonic/gin"
	"go-store/pkg/mytools"
	"go-store/service"
)

func CreateFavorites(context *gin.Context) {
	claim, _ := mytools.ParseToken(context.GetHeader("Authorization"))
	createFavoritesService := service.FavoritesService{}
	if err := context.ShouldBind(&createFavoritesService); err == nil {
		res := createFavoritesService.CreateFavoritesService(context.Request.Context(), claim.ID)
		context.JSON(200, res)
	} else {
		context.JSON(400, MyError(err))
		mytools.Logger.Infoln("err: ", err)
	}
}

func DeleteFavorites(context *gin.Context) {
	claim, _ := mytools.ParseToken(context.GetHeader("Authorization"))
	deleteFavorites := service.FavoritesService{}
	if err := context.ShouldBind(&deleteFavorites); err == nil {
		res := deleteFavorites.DeleteFavorites(context.Request.Context(), claim.ID, context.Param("id"))
		context.JSON(200, res)
	} else {
		context.JSON(400, MyError(err))
		mytools.Logger.Infoln("err: ", err)
	}
}

func ShowFavorites(context *gin.Context) {
	showFavorites := service.FavoritesService{}
	claim, _ := mytools.ParseToken(context.GetHeader("Authorization"))
	if err := context.ShouldBind(&showFavorites); err == nil {
		res := showFavorites.ShowFavorites(context.Request.Context(), claim.ID)
		context.JSON(200, res)
	} else {
		context.JSON(400, MyError(err))
		mytools.Logger.Infoln("err: ", err)
	}
}
