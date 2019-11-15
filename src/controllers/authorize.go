package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	_"Kcoin-Golang/src/models"
)

type AuthoController struct {
	beego.Controller
}

type Json struct{
	Name string `json : "name"`
	Uri string `json:"uri"`
}

func (c * AuthoController) Get(){
	var code string=c.GetString("code")
	text:=models.GetJson(code)

	//存储用户名到cooike中，获取语法：c.Ctx.GetCookie("userName")
	c.Ctx.SetCookie("userName",text.Data.Name,100)
	//存储用户头像url到cooike中，获取语法：c.Ctx.GetCookie("userName")
	c.Ctx.SetCookie("headShotUrl",text.Data.Uri,100)
	//存储用户登录状态到cooike中，其中1表示已登录，获取语法：c.Ctx.GetCookie("userName")
	c.Ctx.SetCookie("status", string('1'),100)

	c.Redirect("/homepage",302)

}

