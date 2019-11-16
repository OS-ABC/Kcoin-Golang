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
	//存储用户名到cooike中，获取语法：c.Ctx.GetCookie("userName")
	c.Ctx.SetCookie("userName"," ",100)
	//存储用户头像url到cooike中，获取语法：c.Ctx.GetCookie("userName")
	c.Ctx.SetCookie("headShotUrl"," ",100)
	//存储用户登录状态到cooike中，其中1表示已登录，获取语法：c.Ctx.GetCookie("userName")
	c.Ctx.SetCookie("status", string('0'),100)
	fmt.Println(("congrat,his is log out "))
	c.Redirect("/homepage",302)

}