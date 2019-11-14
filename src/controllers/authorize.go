package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	_"Kcoin-Golang/src/models"
)

type AuthoController struct {
	beego.Controller
}

type Json struct{
	Name string `json : "name"`
	Uri string `json:"uri"`
}

func (c * AuthoController) Get(){
	var code string=c.GetString("code")
	text:=models.GetJson(code)
	c.Redirect("/homepage",302)

}

