package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
)

type PersonalPageController struct {
	beego.Controller
}

func (c *PersonalPageController) Get() {
	//jsonBuf是一个用于调试的静态json，之后会调用webserver的接口，动态获取。
	jsonBuf :=
		`{
		"errorCode": "0",
		"data":{
			"userName": "DoubleJ",
			"headShotUrl": "../static/img/tx2.png"
		}
	}`

	var user UserInfo//user中存放着json解析后获得的数据。
	errorCode := json.Unmarshal([]byte(jsonBuf), &user)//将jsonBuf的数据解析，然后赋值给user，如果出错会返回对应的errorCode
	if errorCode != nil {//出错了，panic
		fmt.Println("there is an error ,sorry ,please continue debug,haha", errorCode.Error())
	}
	c.Data["user"] = user
	c.TplName = "personalPage.html"//该controller对应的页面
}
