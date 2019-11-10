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
	jsonBuf :=
		`{
		"errorCode": "0",
		"data":{
			"userName": "DoubleJ",
			"headShotUrl": "../static/img/tx2.png"
		}
	}`

	var user UserInfo
	errorCode := json.Unmarshal([]byte(jsonBuf), &user)
	if errorCode != nil {
		fmt.Println("there is an error ,sorry ,please continue debug,haha", errorCode.Error())
	}
	c.Data["user"] = user
	c.TplName = "personalPage.html"
}
