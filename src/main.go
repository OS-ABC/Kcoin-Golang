package main

import (
	"Kcoin-Golang/src/models"
	_ "Kcoin-Golang/src/routers"
	"fmt"

	"github.com/astaxie/beego"

)

func main() {
	fmt.Println(models.GetGithubRepos("Darkone0"))
	beego.Run()
}
