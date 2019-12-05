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

func (c *AuthoController) Get() {
	var code string = c.GetString("code")
	accessToken, _ := getAccessToken(code)
	text := getUserJson(accessToken)

	name := text.Data.Name
	id := text.Data.Id
	// 修改参数
	GithubUser.setGithubUserAccessToken(id,name, accessToken)
	uri := text.Data.Uri
	time = time.Now().Format("2006-01-02 15:04:05.000000")
	
	o := orm.NewOrm()
	o.Using("default")
//  移到model 改成GitID查询
	res, _ := models.FinduserByGitId(id)
	if res == nil {
		insertSql := `INSERT INTO "K_User" (USER_NAME,REGISTER_TIME,HEAD_SHOT_URL,GITHUB_USER_ID) VALUES (?,?,?,?);`
		_, err := o.Raw(insertSql, name, uri, time, id).Exec()
		err := models.InsertUser(name,uri,id)

		if err != nil {
			panic(err)
		}
	}else {
		updateSql := `update "K_User" set register_time = ? where user_name = ?`
		_, err := o.Raw(updateSql, time, name).Exec()
		if err != nil {
			panic(err)
		}
	}


	//存储用户名到cooike中，获取语法：c.Ctx.GetCookie("userName")
	c.Ctx.SetCookie("userName", text.Data.Name, 3600)
	//存储用户头像url到cooike中，获取语法：c.Ctx.GetCookie("userName")
	c.Ctx.SetCookie("headShotUrl", text.Data.Uri, 3600)
	//存储用户登录状态到cooike中，其中1表示已登录，获取语法：c.Ctx.GetCookie("userName")

	c.Ctx.SetCookie("status", string('1'), 3600)

	if redirectUrl := c.Ctx.GetCookie("lastUri"); redirectUrl != "" {
		fmt.Printf(redirectUrl)
		c.Redirect(c.Ctx.GetCookie("lastUri"), 302)
	} else {
		c.Redirect("homepage", 302)
	}
}
