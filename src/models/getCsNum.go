package models

import (
	
	"github.com/astaxie/beego/orm"
)

type Cscount struct {
	user_cs int
}

func GetCsNum(github_id string) interface{} {

	o := orm.NewOrm()
	_ = o.Using("default")

	var list orm.ParamsList
	num,err := o.Raw("select user_cs from `K_User_in_Project` where user_id in (select k_user_id from `K_User` where github_user_id = ?)", github_id).ValuesFlat(&list)
	if err == nil && num > 0{
		fmt.Println(list)
	}
	return list[0]
}
