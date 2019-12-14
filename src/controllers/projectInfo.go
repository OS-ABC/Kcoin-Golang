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
	id := c.GetSession(":id")
	if id == nil {
		id = c.Ctx.Input.Param(":id")
		c.SetSession(":id", id)
	}

	/*id := c.Ctx.Input.Param(":id")*/
	c.Data["id"] = id
	fakeURL := "https://github.com/Darkone0/weatherForcast"

	// starNum := models.GetStarNum(fakeURL)
	// contributorsNum := models.GetContributorNum(fakeURL)
	//这里的session都没有对不同项目进行区分，后续应当还需要更改
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
	projectCC, _ := models.GetProjectsCC(id.(string))
	c.Data["projectCC"] = projectCC
	c.Data["starNum"] = starNum
	c.Data["contributorsNum"] = contributorsNum
	c.TplName = "projectInfo.html"
}
