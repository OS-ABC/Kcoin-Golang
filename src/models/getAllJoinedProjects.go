package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

func GetAllJoinedProjects(userId string) (joinedProjects []Project, err error) {
	var jp []Project
	//var maps []orm.Params
	o := orm.NewOrm()
	_ = o.Using("default")
	SQLQuery := getAllJoinedProjectsQuery()
	if _, err := o.Raw(SQLQuery, userId).QueryRows(&jp); err != nil {
		fmt.Print(err.Error())
		return nil, err
	}

	SQLQuery = getMemberListQuery()
	for _, proj := range jp {
		if _, err = o.Raw(SQLQuery, proj.ProjectId).QueryRows(&proj.MemberList); err != nil {
			fmt.Print(err.Error())
			return nil, err
		}
	}
	return jp, nil

}
func getAllJoinedProjectsQuery() string {
	return "select * from \"K_Project\" where project_id in " +
		"(select project_id from \"K_User_in_Project\" where user_id = ?)"
}
func getMemberListQuery() string {
	return "SELECT user_name FROM \"K_User\" WHERE user_id in " +
		"(SELECT user_id FROM \"K_User_in_Project\" WHERE project_id = ?)"
}

//测试
//func main() {
//	fmt.Print(GetAllJoinedProjects("1" ))
//}