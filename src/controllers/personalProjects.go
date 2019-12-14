package controllers

import (
	"Kcoin-Golang/src/models"
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/astaxie/beego"
)

var memberList_len int     //获取用户github中项目数量
var joinedprojects_len int //参加项目的数量
var ProjectIntro string

type PersonalProjectsController struct {
	beego.Controller
}

func (c *PersonalProjectsController) GetPersonalInfo() {
	name := c.Ctx.GetCookie("userName")
	userid := c.Ctx.GetCookie("userId")
	userBuf, _ := models.GetUserInfo(name)
	projectBuf, _ := models.GetGithubRepos(name)

	status := c.Ctx.GetCookie("status")
	//判断是否登录，如果未登录，登录后跳转到原页面
	c.Ctx.SetCookie("lastUri", c.Ctx.Request.RequestURI)
	if status == "0" || status == "" {
		defer c.Redirect("/login.html", 302)
	}

	user := models.UserInfo{Data: &models.UserData{}}
	var projects models.ProjectInfo
	errorCode := json.Unmarshal([]byte(userBuf), &user)
	errorCode2 := json.Unmarshal([]byte(projectBuf), &projects)

	if errorCode != nil {
		fmt.Println("Oops, there is an error:( please keep debugging.", errorCode.Error())
	}
	if errorCode2 != nil {
		fmt.Println("Oops, there is an error:( please keep debugging.", errorCode.Error())
	}

	c.Data["user"] = user
	c.Data["repos"] = projects
	c.Data["memberList_len"] = strconv.Itoa(len(projects.Data)) //个人项目数量

	//获取已加入项目
	//使用testid=95
	//testid := "9"
	joinedprojects, _ := models.GetAllJoinedProjects(userid)
	//fmt.Print(models.GetAllJoinedProjects("95")
	c.Data["joinedProjects"] = joinedprojects
	c.Data["joinedprojects_len"] = strconv.Itoa(len(joinedprojects))

	//获取管理的项目
	//fmt.Print(user.Data.UserI)
	managedprojects, _ := models.GetAllManagedProjects(userid)
	c.Data["managedprojects_len"] = strconv.Itoa(len(managedprojects))
	c.Data["managedProjects"] = managedprojects
}

func (c *PersonalProjectsController) Post() {
	//var U models.Project
	pUrl := c.GetString("ProjectUrl")     //项目地址
	pName := c.GetString("ProjectName")   //项目名称
	pIntro := c.GetString("ProjectIntro") //项目介绍
	//uploadname := c.GetString("uploadname") //项目封面
	fmt.Println(pUrl, pName, pIntro)

	//提交图片
	f, h, err := c.GetFile("uploadname")
	if err != nil {
		log.Fatal("getfile err ", err)
	}
	defer f.Close()
	fmt.Println(f, h, err)
	ImportProject(pUrl, "../static/img/projectbg.png")

	c.TplName = "personalProjects.html"
	c.Redirect("personalprojects", 302)
}

func (c *PersonalProjectsController) Get() {
	c.GetPersonalInfo()

	c.TplName = "personalProjects.html"
}
