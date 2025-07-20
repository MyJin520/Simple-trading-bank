package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-store/pkg/mytools"
	"go-store/service"
	"net/http"
)

func UserRegister(ctx *gin.Context) {
	var userRegister service.UserService
	err := ctx.ShouldBind(&userRegister)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, MyError(err))
		mytools.Logger.Infoln("err: ", err)

		return
	}
	res := userRegister.Register(ctx.Request.Context())
	ctx.JSON(http.StatusOK, res)
}

func UserLogin(ctx *gin.Context) {
	var userLogin service.UserService
	err := ctx.ShouldBind(&userLogin)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, MyError(err))
		mytools.Logger.Infoln("err: ", err)
		return
	}
	res := userLogin.Login(ctx.Request.Context())
	ctx.JSON(http.StatusOK, res)
}

func UserUpdate(ctx *gin.Context) {
	var userUpdate service.UserService
	// 获取token
	claims, _ := mytools.ParseToken(ctx.GetHeader("Authorization"))
	err := ctx.ShouldBind(&userUpdate)
	if err == nil {
		res := userUpdate.Update(ctx.Request.Context(), claims.ID)
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, MyError(err))
		mytools.Logger.Infoln("err: ", err)
	}
}

func UpAvatar(ctx *gin.Context) {
	// 1. 首先获取上传的文件 - 这是关键问题所在
	file, fileHeader, err := ctx.Request.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "无法获取上传的文件: " + err.Error(),
		})
		return
	}
	defer file.Close() // 确保关闭文件流

	// 2. 检查文件头是否存在
	if fileHeader == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "文件头信息为空",
		})
		return
	}

	// 3. 获取文件信息
	fileSize := fileHeader.Size
	fileName := fileHeader.Filename
	fmt.Println("上传的文件名:", fileName)

	// 4. 验证并解析Token
	token := ctx.GetHeader("Authorization")
	if token == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "缺少认证Token",
		})
		return
	}

	claims, err := mytools.ParseToken(token)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "无效的Token: " + err.Error(),
		})
		return
	}

	// 5. 创建服务对象
	loadAvatar := service.UserService{}

	// 6. 调用服务层方法处理文件
	res := loadAvatar.PostFile(ctx.Request.Context(), claims.ID, file, fileSize)
	ctx.JSON(http.StatusOK, res)
}

// SendEmail 发送邮件
func SendEmail(ctx *gin.Context) {
	var sendEmail service.EmailService
	// 获取token
	claims, _ := mytools.ParseToken(ctx.GetHeader("Authorization"))
	err := ctx.ShouldBind(&sendEmail)
	if err == nil {
		res := sendEmail.SendEmail(ctx.Request.Context(), claims.ID)
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, MyError(err))
		mytools.Logger.Infoln("err: ", err)
	}
}

// VerifyEmail 验证邮件信息
func VerifyEmail(ctx *gin.Context) {
	var verifyEmail service.VerifyEmailService
	err := ctx.ShouldBind(&verifyEmail)
	if err == nil {
		res := verifyEmail.VerifyEmail(ctx.Request.Context(), ctx.GetHeader("Authorization"))
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, MyError(err))
		mytools.Logger.Infoln("err: ", err)
	}
}

// ShowMoney 显示用户余额
func ShowMoney(ctx *gin.Context) {
	var showMoney service.ShowMoneyService
	// 获取token
	claims, _ := mytools.ParseToken(ctx.GetHeader("Authorization"))
	err := ctx.ShouldBind(&showMoney)
	if err == nil {
		res := showMoney.ShowMoney(ctx.Request.Context(), claims.ID)
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, MyError(err))
		mytools.Logger.Infoln("err: ", err)
	}

}
