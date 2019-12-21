package controllers

import (
	"Kcoin-Golang/src/models"
	_ "Kcoin-Golang/src/models"
	// "encoding/json"
	"fmt"

	"github.com/astaxie/beego"
)

type PersonalPageController struct {
	beego.Controller
}

func (c *PersonalPageController) Get() {
	//jsonBuf是一个用于调试的静态json，之后会调用webserver的接口，动态获取。
	//jsonBuf :=
	//	`{
	//	"errorCode": "0",
	//	"data":{
	//		"userName": "DoubleJ",
	//		"headShotUrl": "../static/img/tx2.png"
	//	}
	//}`
	status := c.Ctx.GetCookie("status")
	c.Ctx.SetCookie("lastUri", c.Ctx.Request.RequestURI)
	if status == "0" || status == "" {
		defer c.Redirect("/login.html", 302)
	}

	//获取GitHubId
	//gitId := c.GetSession("GitHubId").(string)
	githubId := c.Ctx.GetCookie("githubId")
	//通过gitHubId查询cs数
	csNum := models.GetCsNum(githubId)

	user := models.UserInfo{Data: &models.UserData{}} //user中存放着json解析后获得的数据。
	user.Data.UserName = c.Ctx.GetCookie("userName")
	user.Data.HeadShotUrl = c.Ctx.GetCookie("headShotUrl")
	user.Data.CsNum = csNum

	c.Data["user"] = user
	c.TplName = "personalPage.html" //该controller对应的页面

	// 函数定义在models目录下的searchCcAndCs.go中，根据用户名查询CC余额
	remainingCc, err := models.GetPersonalRemainingCc(user.Data.UserName)
	if err != nil {
		fmt.Println("you r in personalPage controller, something got wrong "+
			"while querying the database: ", err.Error())
	}

	c.Data["remainingCc"] = remainingCc
}
