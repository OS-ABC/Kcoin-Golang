package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
)

type PersonalProjectsController struct {
	beego.Controller
}

func (c *PersonalProjectsController) Get() {

	jsonBuf :=
		`{
    "errorCode": "0",
    "data": {
        "userName": "Cyan",
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
    }
    }`

	var user UserInfo
	errorCode := json.Unmarshal([]byte(jsonBuf), &user)
	if errorCode != nil {
		fmt.Println("Oops, there is an error:( please keep debugging.", errorCode.Error())
	}

	c.Data["user"] = user
	c.TplName = "personalProjects.html"
}
