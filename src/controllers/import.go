package controllers

import (
	"Kcoin-Golang/src/models"
	"encoding/json"
	"fmt"
	"strconv"
	"github.com/astaxie/beego"
)

type ImportController struct {
	beego.Controller
}

var memberList_len int //获取用户github中项目数量

func (c *ImportController) Get() {
	jsonBuf :=
		`{
    "errorCode": "0",
    "data": {
    	"userName": "anaana",
    	"headshotUrl": "../static/img/tx2.png",
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
    }
    }`
	var user models.UserInfo
	errorCode := json.Unmarshal([]byte(jsonBuf), &user)
	if errorCode != nil {
		fmt.Println("there is an error ,sorry ,please continue debug,haha", errorCode.Error())
	}
	c.Data["user"] = user                                               //json数据解包
	c.Data["memberList_len"] = strconv.Itoa(len(user.Data.ProjectList)) //个人项目数量
	c.TplName = "import.tpl"
}
