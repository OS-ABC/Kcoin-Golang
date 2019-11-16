package controllers


import (

	_ "Kcoin-Golang/src/models"
	"fmt"
	"github.com/astaxie/beego"
)

type LogOutController struct {
	beego.Controller
}
func (c *LogOutController) Get(){

	c.Ctx.SetCookie("userName","",100)
	c.Ctx.SetCookie("headShotUrl","",100)
	c.Ctx.SetCookie("status", string('0'),100)
	fmt.Println(("congrat,his is log out "))
	c.Redirect("/homepage",302)

}