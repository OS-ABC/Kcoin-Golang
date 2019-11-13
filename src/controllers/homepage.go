package controllers

import(
    "github.com/astaxie/beego"
    "encoding/json"
    "fmt"
)

type HomePageController struct {
    beego.Controller
}

//对应Json中数据结构的结构体
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
    Data []Project		        `json:"data"`
}

func (c *HomePageController) Get() {
    //测试数据
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

    //把Json字符串中的数据解析到结构体中
    var proj Result
    err := json.Unmarshal([]byte(jsonBuf),&tmp)
    if err != nil {
        fmt.Println("err=", err)
        return
    }

    //把结构体传到模板当中
    c.Data["Projects"] = proj
    //设置当前登录状态
    c.Data["isLogin"] = true
    //设置Get方法对应展示的模板
    c.TplName = "homePage.tpl"
}
