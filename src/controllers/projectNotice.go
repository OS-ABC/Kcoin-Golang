package controllers

import (
	_ "Kcoin-Golang/src/models"
	"Kcoin-Golang/src/service"

	"github.com/astaxie/beego"
)

type ProjectNoticeController struct {
	beego.Controller
}

func (c *ProjectNoticeController) Get() {
	id := c.Ctx.Input.Param(":id")

	c.Data["id"] = id

	fakeURL := "https://github.com/Darkone0/weatherForcast"

	// starNum := models.GetStarNum(fakeURL)
	// contributorsNum := models.GetContributorNum(fakeURL)
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

	c.Data["starNum"] = starNum
	c.Data["contributorsNum"] = contributorsNum
	c.TplName = "projectNotice.html" //该controller对应的页面
}
