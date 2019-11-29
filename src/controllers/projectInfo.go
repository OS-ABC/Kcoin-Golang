package controllers

import (
	"Kcoin-Golang/src/models"

	"github.com/astaxie/beego"
)

type ProjectInfoController struct {
	beego.Controller
}

func (c *ProjectInfoController) Get() {
	id := c.Ctx.Input.Param(":id")
	c.Data["id"] = id
	fakeURL := "https://github.com/Darkone0/weatherForcast"
	starNum := models.GetStarNum(fakeURL)
	contributorsNum := models.GetContributorNum(fakeURL)
	c.Data["starNum"] = starNum
	c.Data["contributorsNum"] = contributorsNum
	c.TplName = "projectInfo.html"
}
