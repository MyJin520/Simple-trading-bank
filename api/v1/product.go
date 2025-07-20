package v1

import (
	"github.com/gin-gonic/gin"
	"go-store/pkg/mytools"
	"go-store/service"
)

// CreateProduct 创建商品
func CreateProduct(context *gin.Context) {
	form, _ := context.MultipartForm()
	files := form.File["file"]
	claim, _ := mytools.ParseToken(context.GetHeader("Authorization"))
	createProductService := service.ProductService{}
	if err := context.ShouldBind(&createProductService); err == nil {
		res := createProductService.CreateProductService(context.Request.Context(), claim.ID, files)
		context.JSON(200, res)
	} else {
		context.JSON(400, MyError(err))
		mytools.Logger.Infoln("err: ", err)
	}
}

func ListProduct(context *gin.Context) {
	listProduct := service.ProductService{}
	if err := context.ShouldBind(&listProduct); err == nil {
		res := listProduct.ListProduct(context.Request.Context())
		context.JSON(200, res)
	} else {
		context.JSON(400, MyError(err))
		mytools.Logger.Infoln("err: ", err)
	}
}

func FindProduct(context *gin.Context) {
	findProduct := service.ProductService{}
	if err := context.ShouldBind(&findProduct); err == nil {
		res := findProduct.FindProduct(context.Request.Context())
		context.JSON(200, res)
	} else {
		context.JSON(400, MyError(err))
		mytools.Logger.Infoln("err: ", err)
	}
}

func ShowProduct(context *gin.Context) {
	showProduct := service.ProductService{}
	if err := context.ShouldBind(&showProduct); err == nil {
		res := showProduct.ShowProduct(context.Request.Context(), context.Param("id"))
		context.JSON(200, res)
	} else {
		context.JSON(400, MyError(err))
		mytools.Logger.Infoln("err: ", err)
	}
}
func ProductCategory(context *gin.Context) {
	var listCategory service.CategoryService
	if err := context.ShouldBind(&listCategory); err == nil {
		res := listCategory.List(context.Request.Context())
		context.JSON(200, res)
	} else {
		context.JSON(400, err)
	}
}
