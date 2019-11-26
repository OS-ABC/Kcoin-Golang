package controllers

import (
	"Kcoin-Golang/src/models"
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/astaxie/beego"
)

type PersonalProjectsController struct {
	beego.Controller
}

func (c *PersonalProjectsController) Post() {
	var U models.Project
	//var U test_Project

	if error := c.ParseForm(&U); error != nil {
		c.Ctx.WriteString("出错了！")
	}

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

	status := c.Ctx.GetCookie("status")
	{
		c.Ctx.SetCookie("lastUri", c.Ctx.Request.RequestURI)
		if status == "0" || status == "" {
			defer c.Redirect("/login.html", 302)
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
	c.Data["memberList_len"] = strconv.Itoa(len(user.Data.ProjectList)) //个人项目数量
	c.Data["test_projectName"] = U.ProjectName
	c.Data["test_projectUrl"] = U.ProjectUrl

	f, h, err := c.GetFile("uploadname")
	if err != nil {
		log.Fatal("getfile err ", err)
	}
	defer f.Close()
	c.Data["test_filename"] = h.Filename

	c.TplName = "personalProjects.html"

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

	status := c.Ctx.GetCookie("status")
	{
		c.Ctx.SetCookie("lastUri", c.Ctx.Request.RequestURI)
		if status == "0" || status == "" {
			defer c.Redirect("/login.html", 302)
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
	c.Data["memberList_len"] = strconv.Itoa(len(user.Data.ProjectList)) //个人项目数量
	c.TplName = "personalProjects.html"
}
