package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
)

type PersonalPageController struct {
	beego.Controller
}

type UserData struct {
	UserName    string
	HeadShotUrl string
}

type UserInfo struct { //首字母必须大写，要不然写不进去
	Code string `json:"errorCode"`
	Data UserData
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
