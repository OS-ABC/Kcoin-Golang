package controllers

import (
	"Kcoin-Golang/src/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
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
	text:=models.GetGithubAuthJson(code)

	name := text.Data.Name
	uri := text.Data.Uri
	o := orm.NewOrm()
	o.Using("default")

	
	querySql := `select * from "K_User" where USER_NAME = ?`
	res, _ := o.Raw(querySql, name).Exec()
	if  res == nil {
		insertSql := `INSERT INTO "K_User" (USER_NAME,REGISTER_TIME,HEAD_SHOT_URL) VALUES (?,now(),?);`
		_, err := o.Raw(insertSql,name,uri).Exec()

		if err != nil {
			panic(err)
		}
	}
	//存储用户名到cooike中，获取语法：c.Ctx.GetCookie("userName")
	c.Ctx.SetCookie("userName",text.Data.Name,3600)
	//存储用户头像url到cooike中，获取语法：c.Ctx.GetCookie("userName")
	c.Ctx.SetCookie("headShotUrl",text.Data.Uri,3600)
	//存储用户登录状态到cooike中，其中1表示已登录，获取语法：c.Ctx.GetCookie("userName")

	c.Ctx.SetCookie("status", string('1'),3600)

	fmt.Printf(c.Ctx.GetCookie("lastUri"))
	c.Redirect(c.Ctx.GetCookie("lastUri"),302)

}

