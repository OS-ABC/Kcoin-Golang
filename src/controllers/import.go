package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
)

type ImportController struct {
	beego.Controller
}

type UserData struct {
	UserName    string          `json:"userName"`
	HeadShotUrl string          `json:"headshotUrl"`
	ProjectList []Project       //`json:""`
}

type UserInfo struct {
    ErrorCode string              `json:"errorCode"`
    Data UserData              `json:"data"`
}

func (c *ImportController) Get(){
    jsonBuf :=
    `{
    "errorCode": "0",
    "data": [
    	"userName": "Joy",
    	"headshotUrl": "../static/img/tx1.png",
    	"projectList":
    	[
	        {
	            "projectName": "天气预报1",
	            "projectCoverUrl": "../static/img/projectbg.png",
	            "projectUrl": "",
	            "memberList": [
	                {
	                    "userName": "Tony",
	                    "headshotUrl": "../static/img/tx2.png"
	                },
	                {
	                    "userName": "Tony",
	                    "headshotUrl": "../static/img/tx1.png"
	                }
	            ]
	        },
	        {
	            "projectName": "天气预报2",
	            "projectCoverUrl": "../static/img/projectbg.png",
	            "projectUrl": "",
	            "memberList": [
	                {
	                    "userName": "Joy",
	                    "headshotUrl": "../static/img/tx1.png"
	                },
	                {
	                    "userName": "Tony",
	                    "headshotUrl": "../static/img/tx2.png"
	                }
	            ]
	        }
        ]
    ]
    }`
	var user UserInfo
	errorCode := json.Unmarshal([]byte(jsonBuf), &user)
	if errorCode != nil {
		fmt.Println("there is an error ,sorry ,please continue debug,haha", errorCode.Error())
	}
	c.Data["user"] = user
	//c.Data[""]
	c.TplName = "import.tpl"
}

