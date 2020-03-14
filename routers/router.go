package routers

import (
	"Kcoin-Golang/controller"

	"github.com/gin-gonic/gin"
)

func RouterInit() *gin.Engine {
	r := gin.Default()
	// TODO 解决跨域问题
	r.Static("/static", "static")
	// 修改gin中与Vue冲突的模板渲染标签
	r.Delims("{[{", "}]}")
	r.LoadHTMLGlob("templates/*")
	// 渲染HTML页面的路由
	r.GET("/login", controller.Login)
	r.GET("/", controller.Index)

	// github Oauth回调路由
	r.GET("/oauth", controller.OAuth)

	r.GET("/v1/isLogin", controller.IsLogin)

	return r
}
