package routers

import (
	"Kcoin-Golang/controller"
	"Kcoin-Golang/middleware"

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

	// apiv1路由组，里面是需要jwt鉴权才能使用的api
	apiv1 := r.Group("v1")
	apiv1.Use(middleware.JWT())
	{
		apiv1.GET("isLogin", controller.IsLogin)
		//获取用户参与的项目以及用户管理的项目
		//这里建立了一个路由组
		projects := apiv1.Group("/my/projects")
		{
			//参与项目
			projects.GET("/join", controller.GetJoinProjects)
			//管理项目
			projects.GET("/manage", controller.GetManageProjects)
		}

		// 项目操作路由组。获取项目CC记录也可以放在此组里面
		projectsOperation := apiv1.Group("/projects")
		{
			//导入项目
			projectsOperation.POST("/add", controller.AddProject)
		}
	}

	return r
}
