package controllers

import (
	"Kcoin-Golang/src/models"
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/astaxie/beego"
)

var memberList_len int //获取用户github中项目数量
var joinedprojects_len int //参加项目的数量
var ProjectIntro string

type PersonalProjectsController struct {
	beego.Controller
}

func (c *PersonalProjectsController) GetPersonalInfo() {
	name := c.Ctx.GetCookie("userName")
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
	testid := "95"
	joinedprojects, _ := models.GetAllJoinedProjects(testid)
	fmt.Print(models.GetAllJoinedProjects("95"))
	c.Data["joinedProjects"] = joinedprojects
	c.Data["joinedprojects_len"] = strconv.Itoa(len(joinedprojects))
}

func (c *PersonalProjectsController) Post() {
	var U models.Project
	//var U test_Project
	ProjectIntro = c.GetString("ProjectIntro")

	if error := c.ParseForm(&U); error != nil {
		c.Ctx.WriteString("出错了！")
	}

	c.GetPersonalInfo()

	//获取刚刚post的数据
	c.Data["test_projectName"] = U.ProjectName
	c.Data["test_projectUrl"] = U.ProjectUrl

	//textfield
	c.Data["test_ProjectIntro"] = ProjectIntro

	//session获取textfiled
	textfield := c.GetSession("TextField")
	if textfield != nil {
		c.DelSession("TextField")
	}
	c.Data["TextField"] = textfield
	//session获取textfiled
	c.SetSession("TextField", ProjectIntro)

	//提交图片
	f, h, err := c.GetFile("uploadname")
	if err != nil {
		log.Fatal("getfile err ", err)
	}
	defer f.Close()
	c.Data["test_filename"] = h.Filename

	c.TplName = "personalProjects.html"

}

func (c *PersonalProjectsController) Get() {
	c.GetPersonalInfo()

	c.TplName = "personalProjects.html"
}
