package controllers

import (
	"Kcoin-Golang/src/models"
	"Kcoin-Golang/src/service"

	"github.com/astaxie/beego"
)

type ProjectInfoController struct {
	beego.Controller
}

func (c *ProjectInfoController) Get() {
	//session获取id

	id := c.Ctx.Input.Param(":id")
	c.Data["id"] = id
	fakeURL := "https://github.com/Darkone0/weatherForcast"

	// starNum := models.GetStarNum(fakeURL)
	// contributorsNum := models.GetContributorNum(fakeURL)
	//这里的session都没有对不同项目进行区分，后续应当还需要更改
	starNum := c.GetSession(id + "starNum")
	if starNum == nil {
		starNum = service.GetStarNum(fakeURL)
		c.SetSession(id+"starNum", starNum)
	}

	contributorsNum := c.GetSession(id + "contributorsNum")
	if contributorsNum == nil {
		contributorsNum = service.GetContributorNum(fakeURL)
		c.SetSession(id+"contributorsNum", contributorsNum)
	}
	projectCC, _ := models.GetProjectsCC(id)
	c.Data["projectCC"] = projectCC
	c.Data["starNum"] = starNum
	c.Data["contributorsNum"] = contributorsNum
	c.TplName = "projectInfo.html"
}
