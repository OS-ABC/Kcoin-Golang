package controllers

import (
	"Kcoin-Golang/src/models"
	_ "Kcoin-Golang/src/models"
	"Kcoin-Golang/src/service"

	"github.com/astaxie/beego"
)

type ProjectMemberListController struct {
	beego.Controller
}

func (c *ProjectMemberListController) Get() {
	//id := c.Ctx.Input.Param(":id")
	id := c.GetSession(":id")
	if id == nil {
		id = c.Ctx.Input.Param(":id")
		c.SetSession(":id", id)
	}
	c.Data["id"] = id
	//解决了session造成的bug后，通过读取项目id返回所有项目的信息
	membersInfo, _ := models.GetMembersInfoByProjectName(id.(string))
	c.Data["membersInfo"] = membersInfo

	fakeURL := "https://github.com/Darkone0/weatherForcast"

	// starNum := models.GetStarNum(fakeURL)
	// contributorsNum := models.GetContributorNum(fakeURL)
	starNum := c.GetSession("starNum")
	if starNum == nil {
		starNum = service.GetStarNum(fakeURL)
		c.SetSession("starNum", starNum)
	}

	contributorsNum := c.GetSession("contributorsNum")
	if contributorsNum == nil {
		contributorsNum = service.GetContributorNum(fakeURL)
		c.SetSession("contributorsNum", contributorsNum)
	}

	c.Data["starNum"] = starNum
	c.Data["contributorsNum"] = contributorsNum

	c.TplName = "projectMemberList.html" //该controller对应的页面
}
