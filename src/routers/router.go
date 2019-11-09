package routers

import (
	"github.com/astaxie/beego"
	"kcoin-golang/src/controllers"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
