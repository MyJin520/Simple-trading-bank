package v1

import (
	"github.com/gin-gonic/gin"
	"go-store/pkg/mytools"
	"go-store/service"
	"net/http"
)

func OrderPlay(ctx *gin.Context) {
	orderPay := service.OrderService{}
	claim, err := mytools.ParseToken(ctx.GetHeader("Authorization"))
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	if err := ctx.ShouldBind(&orderPay); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res := orderPay.PayDown(ctx.Request.Context(), claim.ID, ctx.Param("id"))
	ctx.JSON(http.StatusOK, res)
}
