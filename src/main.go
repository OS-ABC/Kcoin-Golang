package main

import (
	_ "Kcoin-Golang/src/routers"
	"github.com/astaxie/beego"
)

func main() {
	//fmt.Println(models.GetGithubRepos("Darkone0"))
	//var users []string
	//users=append(users, "scarydemon2","Darkone0")
	//models.SendEMailToPotentialUsers(users)
	beego.Run()
}
