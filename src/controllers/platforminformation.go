package controllers

import (
	"Kcoin-Golang/src/models"
	_ "Kcoin-Golang/src/models"
	"fmt"
	"github.com/astaxie/beego"
)

type PlatformInformationController struct {
	beego.Controller
}

func (c *PlatformInformationController) Get() {
	c.TplName = "platformInformation.html"		//该controller对应的页面

	platfmProjNum, err := models.PlatfmProjNum()
	if err != nil {
		fmt.Println("Oops, something's wrong: ", err.Error())
	}
	c.Data["projNum"] = platfmProjNum
}