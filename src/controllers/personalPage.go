package controllers

import (
	"Kcoin-Golang/src/models"
	_ "Kcoin-Golang/src/models"
	"encoding/json"
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
	status:=c.Ctx.GetCookie("status")
	{
		c.Ctx.SetCookie("lastUri",c.Ctx.Request.RequestURI)
		if status =="0"||status ==""{
			defer c.Redirect("/login.html",302)
		}
	}
	userName := c.Ctx.GetCookie("userName")
	user := models.UserInfo{Data:&models.UserData{}}//user中存放着json解析后获得的数据。
	jsonBuf , _ := models.GetUserInfo(userName)
	errorCode := json.Unmarshal([]byte(jsonBuf), &user)//将jsonBuf的数据解析，然后赋值给user，如果出错会返回对应的errorCode
	if errorCode != nil {//出错了，panic
		fmt.Println("you r in personalPage controller ,there is ia bug ,and the information is : ", errorCode.Error())
	}
	c.Data["user"] = user
	c.TplName = "personalPage.html"//该controller对应的页面
}
