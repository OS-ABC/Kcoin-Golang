package routers

import (
	"Kcoin-Golang/src/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/homepage", &controllers.HomePageController{})
	beego.Router("/logout", &controllers.LogOutController{})
	beego.Router("/personalpage", &controllers.PersonalPageController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/join", &controllers.JoinController{})
	beego.Router("/autho", &controllers.AuthoController{})
	beego.Router("/import", &controllers.ImportController{})
	beego.Router("/personalprojects", &controllers.PersonalProjectsController{})
	beego.Router("/projectInfo", &controllers.ProjectInfoController{})
	beego.Router("/projectMemberList", &controllers.ProjectMemberListController{})
	beego.Router("/projectNotice", &controllers.ProjectNoticeController{})
}
