package routes

import (
	"github.com/gin-gonic/gin"
	api "go-store/api/v1"
	"go-store/middleware"
	"net/http"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.CorsMiddle())
	r.StaticFS("/static", http.Dir("./static"))

	v1 := r.Group("api/v1")
	{
		v1.GET("ping", func(c *gin.Context) {
			c.JSON(200, "success")
		})
		v1.POST("user/register", api.UserRegister) // 用户注册
		v1.POST("user/user-login", api.UserLogin)  // 用户登录
		v1.GET("carousel", api.ListCarousel)       // 轮播图
		v1.GET("product", api.ListProduct)
		v1.POST("find-product", api.FindProduct)       // 搜索商品
		v1.GET("show-product/:id", api.ShowProduct)    // 展示商品
		v1.GET("productCategory", api.ProductCategory) // 商品分类

		// 需要保护的路由组
		authed := v1.Group("/")
		authed.Use(middleware.JWT())
		{
			// 用户操作模块
			authed.POST("user", api.UserUpdate)
			authed.POST("load-avatar", api.UpAvatar)
			authed.POST("send-email", api.SendEmail)
			authed.POST("verify-email", api.VerifyEmail)
			authed.POST("show-money", api.ShowMoney)

			// 商品操作模块
			authed.POST("product", api.CreateProduct)

			// 收藏操作模块
			authed.GET("show-favorites", api.ShowFavorites)
			authed.POST("create-favorites", api.CreateFavorites)
			authed.DELETE("favorites/:id", api.DeleteFavorites)

			// 地址模块
			authed.GET("show-address", api.ShowAddress)
			authed.POST("create-address", api.CreateAddress)
			authed.DELETE("delete-address/:id", api.DeleteAddress)
			authed.POST("update-address/:id", api.UpdateAddress)
			authed.GET("get-address/:id", api.GetByIDAddress)

			// 购物车模块
			authed.GET("show-shopping", api.ShowShoppingCart)
			authed.POST("create-shopping", api.CreateShoppingCart)
			authed.DELETE("delete-shopping/:id", api.DeleteShoppingCart)
			authed.PUT("update-shopping/:id", api.UpdateShoppingCart)

			// 订单模块
			authed.POST("create-order", api.CreateOrder)         // 创建订单接口，接收必要参数后创建新订单
			authed.GET("show-order", api.ShowOrder)              // 展示所有订单接口，返回当前用户的所有订单信息
			authed.GET("show-byID-order/:id", api.ShowOrderByID) // 根据订单 ID 展示单个订单接口，返回指定 ID 的订单详细信息
			authed.DELETE("delete-order/:id", api.DeleteOrder)   // 根据订单 ID 删除订单接口，删除指定 ID 的订单记录

			// 支付功能
			authed.POST("playDown/:id", api.OrderPlay)

		}
	}

	return r
}
