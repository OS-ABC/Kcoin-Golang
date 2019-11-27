package controllers

import (
	_ "Kcoin-Golang/src/models"

	"github.com/astaxie/beego"
)

type ProjectFundingController struct {
	beego.Controller
}

func (c *ProjectFundingController) Get() {
	c.TplName = "projectFunding.html"		//该controller对应的页面
}