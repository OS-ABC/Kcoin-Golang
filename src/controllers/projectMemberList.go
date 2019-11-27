package controllers

import (
	_ "Kcoin-Golang/src/models"

	"github.com/astaxie/beego"
)

type ProjectMemberListController struct {
	beego.Controller
}

func (c *ProjectMemberListController) Get() {
	id := c.Ctx.Input.Param(":id")
	c.Data["id"] = id
	c.TplName = "projectMemberList.html" //该controller对应的页面
}
