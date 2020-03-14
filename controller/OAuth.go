package controller

import (
	"Kcoin-Golang/models"
	"Kcoin-Golang/service"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// 渲染登录html页面
func Login(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

// github回调的OAuth地址
// TODO 所有跟操作数据库有关的语句应该封装到models包中
func OAuth(c *gin.Context) {
	code := c.Query("code")
	accessToken, err := service.GetAccessToken(code)

	if err != nil {
		fmt.Println("get access token failed!, err: ", err)
	}
	var user *models.User
	user, err = service.GetGithubUserInfo(accessToken)
	if err != nil {
		fmt.Println("Get github user info failed, err: ", err)
	}
	// 把accesstoken存储到map中
	service.GithubAccessToken.Store(user.GithubID, accessToken)
	temUser := &models.User{}
	models.DB.Where("github_id = ?", user.GithubID).First(temUser)
	// 数据库中没有这条用户记录就插入，有就更新
	fmt.Println(temUser)
	if temUser.Name == "" {
		user.Time = time.Now()
		models.DB.Debug().Create(user)
	} else {
		models.DB.Debug().Where("github_id = ?", user.GithubID).Update(user)
	}
	// TODO 使用jwt来记住用户和鉴定用户权限
	c.SetCookie("jwt", "test", 3600, "/", "localhost", false, true)
	// TODO 跳转回刚才访问的页面或者首页
	c.Redirect(302, "/")
}
