package middleware

import (
	"github.com/gin-gonic/gin"
	"go-store/pkg/e"
	"go-store/pkg/mytools"
	"net/http"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		code := http.StatusOK
		token := c.GetHeader("Authorization")
		if token == "" {
			code = http.StatusNotFound
		} else {
			claims, err := mytools.ParseToken(token)
			if err != nil {
				code = 30007 // 获取token失败
			} else if time.Now().Unix() > claims.ExpiresAt.Unix() {
				code = 30008 // token过期
			}
		}
		if code != http.StatusOK {
			c.JSON(200, gin.H{
				"status": code,
				"msg":    e.GetMSG(code),
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
