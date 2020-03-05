// TODO Project相关的controller可以全部放到这个文件下, 即这个文件有若干Controller.
package controllers

import (
	_ "Kcoin-Golang/src/models"

	"github.com/astaxie/beego"
)

type ProjectFundingController struct {
	beego.Controller
}

func (c *ProjectFundingController) Get() {
	c.TplName = "projectFunding.html" //该controller对应的页面
}
