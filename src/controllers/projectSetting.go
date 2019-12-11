package controllers

import (
	"Kcoin-Golang/src/models"
	_ "Kcoin-Golang/src/models"

	"github.com/astaxie/beego"
)

type ProjectSettingController struct {
	beego.Controller
}

func (c *ProjectSettingController) Get() {
	//id := c.Ctx.Input.Param(":id")
	id := c.GetSession(":id")
	if id == nil {
		id = c.Ctx.Input.Param(":id")
		c.SetSession(":id", id)
	}
	c.Data["id"] = id

	fakeURL := "https://github.com/Darkone0/weatherForcast"

	// starNum := models.GetStarNum(fakeURL)
	// contributorsNum := models.GetContributorNum(fakeURL)
	starNum := c.GetSession("starNum")
	if starNum == nil {
		starNum = models.GetStarNum(fakeURL)
		c.SetSession("starNum", starNum)
	}

	contributorsNum := c.GetSession("contributorsNum")
	if contributorsNum == nil {
		contributorsNum = models.GetContributorNum(fakeURL)
		c.SetSession("contributorsNum", contributorsNum)
	}

	c.Data["starNum"] = starNum
	c.Data["contributorsNum"] = contributorsNum
	c.TplName = "projectSetting.html" //该controller对应的页面

	//session获取textfiled
	textfield := c.GetSession("TextField")
	if textfield != nil {
		c.Data["TextField"] = textfield
	} else {
		c.Data["TextField"] = "不超过200字符"
	}

}
