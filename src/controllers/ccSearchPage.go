package controllers

import (
	"Kcoin-Golang/src/models"
	_ "Kcoin-Golang/src/models"
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
)

type CcSearchPageController struct {
	beego.Controller
}

func (c *CcSearchPageController) Get() {
	//jsonBuf是一个用于调试的静态json,之后从数据库中读取数据
	jsonBuf :=
		`{
		"errorCode": "0",
		"data": {
			"userName": "Sonorous ",
			"headshotUrl": "../static/img/tx2.png",
			"userCcOpeList":
			[
				{
					"opeCcDate":"11.18",
					"opeCcType":"sell",
					"opeCcNumber":"-2000"
				},
				{
					"opeCcDate":"11.18",
					"opeCcType":"buy",
					"opeCcNumber":"5000"
				},
				{
					"opeCcDate":"11.19",
					"opeCcType":"sell",
					"opeCcNumber":"-1000"
				},
				{
					"opeCcDate":"11.18",
					"opeCcType":"sell",
					"opeCcNumber":"-1000"
				},
				{
					"opeCcDate":"11.18",
					"opeCcType":"buy",
					"opeCcNumber":"1000"
				},
				{
					"opeCcDate":"11.19",
					"opeCcType":"sell",
					"opeCcNumber":"-500"
				},
				{
					"opeCcDate":"11.18",
					"opeCcType":"sell",
					"opeCcNumber":"-3000"
				},
				{
					"opeCcDate":"11.18",
					"opeCcType":"buy",
					"opeCcNumber":"2000"
				},
				{
					"opeCcDate":"11.19",
					"opeCcType":"sell",
					"opeCcNumber":"-800"
				},
				{
					"opeCcDate":"11.18",
					"opeCcType":"sell",
					"opeCcNumber":"-1000"
				}
			]
		}
	}`

	status:=c.Ctx.GetCookie("status")
	{
		c.Ctx.SetCookie("lastUri",c.Ctx.Request.RequestURI)
		if status =="0"||status ==""{
			defer c.Redirect("/login.html",302)
		}
	}
	var user models.UserInfo
	errorCode := json.Unmarshal([]byte(jsonBuf), &user)
	user.Data.UserName = c.Ctx.GetCookie("userName")
	user.Data.HeadShotUrl = c.Ctx.GetCookie("headShotUrl")
	if errorCode != nil {
		fmt.Println("Oops, there is an error:( please keep debugging.", errorCode.Error())
	}

	c.Data["user"] = user
	c.TplName = "ccSearchPage.html"//该controller对应的页面
}
