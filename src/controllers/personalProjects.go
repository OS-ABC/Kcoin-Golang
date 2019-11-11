package controllers

import(
	"github.com/astaxie/beego"
	"encoding/json"
	"fmt"
)

type PersonalProjectsController struct {
    beego.Controller
}

func (c *PersonalProjectsController) Get() {

    jsonBuf :=
    `{
    "errorCode": "0",
    "data": [
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
    }`

    var user UserInfo
    var proj Result
    errorCode1 := json.Unmarshal([]byte(jsonBuf), &user)
    errorCode2 := json.Unmarshal([]byte(jsonBuf), &proj)
	if errorCode1 != nil {
		fmt.Println("Oops, there is an error:( please keep debugging.", errorCode1.Error())
    }
	if errorCode2 != nil {
		fmt.Println("Oops, there is an error:( please keep debugging.", errorCode2.Error())
    }

	c.Data["user"] = user
	c.Data["Projects"] = proj
	c.TplName = "personalProjects.html"
}
