package controllers

import (
	"github.com/astaxie/beego"
)

type ProjectInfoController struct {
	beego.Controller
}

func (c *ProjectInfoController) Get() {
	c.TplName = "projectInfo.html"
}