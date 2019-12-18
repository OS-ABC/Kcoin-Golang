package models

import (
	"encoding/json"
	_ "encoding/json"
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
	return `select user_cs, b.user_cc from "k_user_in_project" a, "k_user" b where a.user_id=? and a.project_id=? and a.user_id = b.user_id`
}

// 实现personalPage控制器中的查询CC余额
func GetPersonalRemainingCc(userName string) (string, error) {
	o1 := orm.NewOrm()
	o1.Using("default")
	ccQuery := `select user_cc from "k_user" a where a.user_name=?`
	// 分别代表数据库中查到的余额和错误
	sum, err1 := o1.Raw(ccQuery, userName).Exec()
	if err1 != nil {
		fmt.Println(err1.Error())
		return "", err1
	}
	// json.Marshal函数将sum封装成json格式存进res，同时返回错误信息
	res, err2 := json.Marshal(sum)
	if err2 != nil {
		fmt.Println(err2.Error())
		return "", err2
	}
	return string(res), nil
}
