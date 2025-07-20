package v1

import (
	"github.com/gin-gonic/gin"
	"go-store/service"
)

func ListCarousel(context *gin.Context) {
	var ListCarousel service.CarouselService
	if err := context.ShouldBind(&ListCarousel); err == nil {
		res := ListCarousel.List(context.Request.Context())
		context.JSON(200, res)
	} else {
		context.JSON(400, err)
	}
}
