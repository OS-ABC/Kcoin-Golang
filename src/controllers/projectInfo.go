package controllers

import (
	"github.com/astaxie/beego"
)

type ProjectInfoController struct {
	beego.Controller
}

func (c *ProjectInfoController) Get() {
	id := c.Ctx.Input.Param(":id")
	c.Data["id"] = id
	c.TplName = "projectInfo.html"
}
