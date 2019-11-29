package controllers

import (
	_ "Kcoin-Golang/src/models"

	"github.com/astaxie/beego"
)

type PlatformInformationController struct {
	beego.Controller
}

func (c *PlatformInformationController) Get() {
	c.TplName = "platformInformation.html"		//该controller对应的页面
}