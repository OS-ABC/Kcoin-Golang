package routers

import (
	"Kcoin-Golang/controller"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func RouterInit() *gin.Engine {
	r := gin.Default()
	// cors中间件解决跨域问题
	r.Use(cors.Default())
	r.Static("/static", "static")
	// 修改gin中与Vue冲突的模板渲染标签
	r.Delims("{[{", "}]}")
	// 加载html页面
	r.LoadHTMLGlob("templates/*")
	// 用于渲染HTML页面的路由
	r.GET("/login", controller.Login)
	r.GET("/", controller.Index)

	// github Oauth回调路由
	r.GET("/oauth", controller.OAuth)
	// 判断用户是否已经登录
	r.GET("/v1/isLogin", controller.IsLogin)

	return r
}
