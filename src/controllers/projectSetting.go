package controllers

import (
	_ "Kcoin-Golang/src/models"

	"github.com/astaxie/beego"
)

type ProjectSettingController struct {
	beego.Controller
}

func (c *ProjectSettingController) Get() {
	id := c.Ctx.Input.Param(":id")
	c.Data["id"] = id
	c.TplName = "projectSetting.html" //该controller对应的页面

	//session获取textfiled
	textfield := c.GetSession("TextField")
	if textfield != nil {
		c.Data["TextField"] = textfield
	} else {
		c.Data["TextField"] = "不超过200字符"
	}

}
