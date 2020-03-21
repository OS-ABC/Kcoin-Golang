package middleware

import (
	"Kcoin-Golang/service"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// TODO 返回的json中应该有个code字段，表示错误码
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var message string
		message = "ok"
		token, _ := c.Cookie("jwt")
		if token == "" {
			message = "NO_TOKEN"
		} else {
			claims, err := service.ParseToken(token)
			if err != nil {
				message = "ERROR_AUTH_CHECK_TOKEN_FAIL"
			} else if time.Now().Unix() > claims.ExpiresAt {
				message = "ERROR_AUTH_CHECK_TOKEN_TIMEOUT"
			}
		}

		if message != "ok" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg":   message,
				"token": token,
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
