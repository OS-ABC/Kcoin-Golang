package main

import (
	_ "Kcoin-Golang/src/routers"
	"github.com/astaxie/beego"
)

func main() {

	//fmt.Println(models.GetGithubRepos("Darkone0"))
	beego.Run()
}
