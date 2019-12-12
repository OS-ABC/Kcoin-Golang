package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

func GetCcAndCsQuery(userId int, projectId int)([]orm.Params, error){
	o := orm.NewOrm()
	_ = o.Using("default")
	SQLQuery := getQuery()
	var maps []orm.Params
	if _, err := o.Raw(SQLQuery, userId, projectId).Values(&maps); err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	if len(maps) == 0 {
		return nil, new(noResultErr)
	}
	return maps, nil
}

func getQuery() string {
	return `select user_cs, b.user_cc from "K_User_in_Project" a, "K_User" b where a.user_id=? and a.project_id=? and a.user_id = b.user_id`
}


