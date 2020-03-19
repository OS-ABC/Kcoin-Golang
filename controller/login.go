package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func IsLogin(c *gin.Context) {
	jwt, err := c.Cookie("jwt")
	if err != nil {
		fmt.Println("get jwt cookie failed, err: ", err)
	}
	// TODO 加上jwt的验证，现在是假的
	if jwt == "" {
		c.JSON(http.StatusOK, gin.H{"isLogin": false})
	} else {
		// TODO 根据jwt获取信息并传回前端
		c.JSON(http.StatusOK, gin.H{"isLogin": true, "userName": "hahaha"})
	}
}
