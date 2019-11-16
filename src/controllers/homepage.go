package controllers

import (
    "Kcoin-Golang/src/models"
    "encoding/json"
    "fmt"
    "github.com/astaxie/beego"
)

type HomePageController struct {
    beego.Controller
}


func (c *HomePageController) Get() {
    //测试数据
    //jsonBuf :=
    //`{
    //"errorCode": "0",
    //"data": [
    //    {
    //        "projectName": "天气预报1",
    //        "projectCoverUrl": "../static/img/projectbg.png",
    //        "projectUrl": "",
    //        "memberList": [
    //            {
    //                "userName": "Tony",
    //                "headshotUrl": "../static/img/tx2.png"
    //            },
    //            {
    //                "userName": "Tony",
    //                "headshotUrl": "../static/img/tx1.png"
    //            }
    //        ]
    //    },
    //    {
    //        "projectName": "天气预报1",
    //        "projectCoverUrl": "../static/img/projectbg.png",
    //        "projectUrl": "",
    //        "memberList": [
    //            {
    //                "userName": "Tony",
    //                "headshotUrl": "../static/img/tx2.png"
    //            },
    //            {
    //                "userName": "Tony",
    //                "headshotUrl": "../static/img/tx1.png"
    //            }
    //        ]
    //    },
    //    {
    //        "projectName": "天气预报1",
    //        "projectCoverUrl": "../static/img/projectbg.png",
    //        "projectUrl": "",
    //        "memberList": [
    //            {
    //                "userName": "Tony",
    //                "headshotUrl": "../static/img/tx2.png"
    //            },
    //            {
    //                "userName": "Tony",
    //                "headshotUrl": "../static/img/tx1.png"
    //            }
    //        ]
    //    },
    //    {
    //        "projectName": "天气预报1",
    //        "projectCoverUrl": "../static/img/projectbg.png",
    //        "projectUrl": "",
    //        "memberList": [
    //            {
    //                "userName": "Tony",
    //                "headshotUrl": "../static/img/tx2.png"
    //            },
    //            {
    //                "userName": "Tony",
    //                "headshotUrl": "../static/img/tx1.png"
    //            }
    //        ]
    //    },
    //    {
    //        "projectName": "天气预报1",
    //        "projectCoverUrl": "../static/img/projectbg.png",
    //        "projectUrl": "",
    //        "memberList": [
    //            {
    //                "userName": "Tony",
    //                "headshotUrl": "../static/img/tx2.png"
    //            },
    //            {
    //                "userName": "Tony",
    //                "headshotUrl": "../static/img/tx1.png"
    //            }
    //        ]
    //    },
    //    {
    //        "projectName": "天气预报2",
    //        "projectCoverUrl": "../static/img/projectbg.png",
    //        "projectUrl": "",
    //        "memberList": [
    //            {
    //                "userName": "Joy",
    //                "headshotUrl": "../static/img/tx1.png"
    //            },
    //            {
    //                "userName": "Tony",
    //                "headshotUrl": "../static/img/tx2.png"
    //            }
    //        ]
    //    }
    //]
    //}`

    //从cookie中取出状态
    isLogin :=c.Ctx.GetCookie("status")


    //把Json字符串中的数据解析到结构体中
    var proj ProjectListInfo
    projectBuf , _ := models.GetAllProjectsInfo()
    err := json.Unmarshal([]byte(projectBuf),&proj)
    if err != nil {
        fmt.Println("err=", err)
        //return
    }


    //把结构体传到模板当中
    c.Data["Projects"] = proj

    //判断当前登录状态
    if isLogin == "1"{
        c.Data["isLogin"] = true
        var user UserInfo//user中存放着json解析后获得的数据。
        user.Data.UserName = c.Ctx.GetCookie("userName")
        user.Data.HeadShotUrl = c.Ctx.GetCookie("headShotUrl")
        c.Data["user"] = user
    } else {
        c.Data["isLogin"] = false
    }


    //设置Get方法对应展示的模板
    c.TplName = "homePage.tpl"
}
