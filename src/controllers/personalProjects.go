package controllers

import (
	"Kcoin-Golang/src/models"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
)

var memberList_len int //获取用户github中项目数量
var ProjectIntro string

type PersonalProjectsController struct {
	beego.Controller
}

//导入项目
func (c *PersonalProjectsController) Post() {
	pUrl := c.GetString("ProjectUrl")       //项目地址
	pName := c.GetString("ProjectName")     //项目名称
	pIntro := c.GetString("ProjectIntro")   //项目介绍
	uploadname := c.GetString("uploadname") //项目封面
	fmt.Println(pUrl, pName, pIntro, uploadname)
	//检查Url合法性
	err := CheckGithubRepoUrl("0", pUrl) //*************写这个方法的人帮忙看一下这里第一个参数id要传什么*********
	if err != nil {
		//url不合法，返回错误并给出提示
	}
	//获取项目所有的开发者信息
	//err = GetAllContributor(pUrl)
	if err != nil {
		//url不合法，后台log输出日志
	}
	//将开发人员信息存入数据库的临时用户表中

	//向被邀请的人发送邮件
	//err = SendEMailToPotentialUsers()

	//返回导入成功信息
}

func (c *PersonalProjectsController) Get() {
	name := c.Ctx.GetCookie("userName")
	projectBuf, _ := models.GetGithubRepos(name)

	status := c.Ctx.GetCookie("status")
	//判断是否登录，如果未登录，登录后跳转到原页面
	c.Ctx.SetCookie("lastUri", c.Ctx.Request.RequestURI)
	if status == "0" || status == "" {
		defer c.Redirect("/login.html", 302)
	}

	user := models.UserInfo{Data: &models.UserData{}} //user中存放着json解析后获得的数据。
	user.Data.UserName = c.Ctx.GetCookie("userName")
	user.Data.HeadShotUrl = c.Ctx.GetCookie("headShotUrl")

	var projects models.ProjectInfo
	errorCode := json.Unmarshal([]byte(projectBuf), &projects)

	if errorCode != nil {
		fmt.Println("Oops, there is an error:( please keep debugging.", errorCode.Error())
	}

	c.Data["user"] = user
	c.Data["repos"] = projects
	c.Data["memberList_len"] = strconv.Itoa(len(projects.Data)) //个人项目数量
	c.TplName = "personalProjects.html"

}
