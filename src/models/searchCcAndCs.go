package models

import (
	// "encoding/json"
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
func GetPersonalRemainingCc(userName string) (float64, error) {
	o := orm.NewOrm()
	_ = o.Using("default")

	type UsrCC struct {
		User_name    string
		User_cc      float64
	}
	var userCc UsrCC

	/**将取回的用户名和CC余额赋值给结构体userCc; 如有错误，则赋值给err1
	 ***结构体中属性名需要与数据库中对应字段相同，且首字母大写*/
	ccQuery := `SELECT user_name, user_cc FROM "K_User" WHERE USER_NAME=?`
	err1 := o.Raw(ccQuery, userName).QueryRow(&userCc)
	if err1 != nil {
		fmt.Println(err1.Error())
		return -1.0, err1
	}

	return userCc.User_cc, nil

}