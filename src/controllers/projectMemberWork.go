package controllers

import (
	_ "Kcoin-Golang/src/models"

	"github.com/astaxie/beego"
)

type ProjectMemberWorkController struct {
	beego.Controller
}

func (c *ProjectMemberWorkController) Get() {
	id := c.Ctx.Input.Param(":id")
	c.Data["id"] = id
	c.TplName = "projectMemberWork.html" //该controller对应的页面
}
