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
    //从cookie中取出状态
    isLogin :=c.Ctx.GetCookie("status")

    //把Json字符串中的数据解析到结构体中
    var proj models.ProjectInfo
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
        user := models.UserInfo{Data:&models.UserData{}}//user中存放着json解析后获得的数据。
        user.Data.UserName = c.Ctx.GetCookie("userName")
        user.Data.HeadShotUrl = c.Ctx.GetCookie("headShotUrl")
        c.Data["user"] = user
    } else {
        c.Data["isLogin"] = false
        c.Ctx.SetCookie("userName","",100)
        c.Ctx.SetCookie("headShotUrl","",100)
        c.Ctx.SetCookie("status", string('0'),100)
        c.Ctx.SetCookie("lastUri",c.Ctx.Request.RequestURI)
    }

    //设置Get方法对应展示的模板
    c.TplName = "homePage.tpl"
}
