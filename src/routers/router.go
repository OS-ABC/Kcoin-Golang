package routers

import (
	"Kcoin-Goloang/src/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/homepage", &controllers.HomePageController{})
	beego.Router("/personalpage", &controllers.PersonalPageController{})
  beego.Router("/login", &controllers.LoginController{})
  beego.Router("/join", &controllers.JoinController{})
  beego.Router("/import",&controllers.ImportController{})
}
