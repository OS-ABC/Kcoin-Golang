// TODO 该文件并入k_user.go中
package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

func GetAllJoinedProjects(userId string) (joinedProjects []*Project, err error) {
	var jp []*Project
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
		var memberlist []*UserData
		if _, err = o.Raw(SQLQuery, proj.ProjectId).QueryRows(&memberlist); err != nil {
			fmt.Print(err.Error())
			return nil, err
		}

		proj.MemberList = memberlist
		fmt.Println(proj.MemberList)
	}
	return jp, nil
}

func getAllJoinedProjectsQuery() string {
	return "select * from \"k_project\" where project_id in " +
		"(select project_id from \"k_user_in_project\" where user_id = ?)"
}

func getMemberListQuery() string {
	return `SELECT u.k_user_id, u.user_name, u.head_shot_url
			FROM "k_user" u LEFT JOIN "k_user_in_project" up on u.k_user_id = up.user_id 
			WHERE up.project_id = ?`
	//return "SELECT user_id, user_name, head_shot_url FROM \"k_user\" WHERE user_id in " +
	//	"(SELECT user_id FROM \"k_user_in_project\" WHERE project_id = ?)"
}

//测试
//func main() {
//	fmt.Print(GetAllJoinedProjects("1" ))
//}
