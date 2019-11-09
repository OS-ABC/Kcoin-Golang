package controllers

import(
    "github.com/astaxie/beego"
    "encoding/json"
    "fmt"
)

type HomePageController struct {
    beego.Controller
}

var projectName string = "天气预报"

type ProjectMember struct {
    MemberName string           `json:"userName"`
    MemberHeadshotUrl string    `json:"headshotUrl"`
}

type Project struct {
    ProjectName string          `json:"projectName"`
    ProjectCoverUrl string      `json:"projectCoverUrl"`
    ProjectUrl string           `json:"projectUrl"`
    MemberList []ProjectMember  `json:"memberList"`
}
type Result struct {
    ErrorCode string            `json:"errorCode"`
    Data []Project		       `json:"data"`
}

func (c *HomePageController) Get() {

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

    var tmp Result
    err := json.Unmarshal([]byte(jsonBuf),&tmp)
    if err != nil {
        fmt.Println("err=", err)
        return
    }
    proj := tmp
    //(*proj).ProjectName = "天气预报233"
    //(*proj).ProjectCoverUrl = "../static/img/tx1.png"

    c.Data["Projects"] = proj
    c.TplName = "HomePage.tpl"
}
