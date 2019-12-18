package models

import (
	"github.com/astaxie/beego/orm"
	"strconv"
)


func GetCsNum(github_id string) int {

	//db, err := sql.Open("postgres", "user=sspkukcoin password=kcoin2019 dbname=postgres host=114.115.133.140 port=5432 sslmode=disable")

	o := orm.NewOrm()
	_ = o.Using("default")

	querySql := `select user_cs from "k_user_in_project" where user_id in (select k_user_id from "k_user" where github_user_id = ?)`
	var maps []orm.Params
	_,err := o.Raw(querySql, github_id).Values(&maps)
	checkErr(err)
	var user_cs string
	for _, term := range maps {
		user_cs = term["user_cs"].(string)
	}
	res, _ := strconv.Atoi(user_cs)
	return res
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
