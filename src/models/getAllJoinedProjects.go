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

func GetAllManagedProjects(userId string) (joinedProjects []*Project, err error) {
	var jp []*Project
	//var maps []orm.Params
	o := orm.NewOrm()
	_ = o.Using("default")
	//获取管理项目
	SQLQuery := getAllManagedProjectsQuery()
	if _, err := o.Raw(SQLQuery, userId).QueryRows(&jp); err != nil {
		fmt.Print(err.Error())
		return nil, err
	}

	//获取项目对应的成员列表
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

//查询全部参加的项目
func getAllJoinedProjectsQuery() string {
	return "select * from \"K_Project\" where project_id in " +
		"(select project_id from \"K_User_in_Project\" where user_id = ?)"
}

//查询子成员项目列表
func getMemberListQuery() string {
	return `SELECT u.k_user_id, u.user_name, u.head_shot_url
			FROM "K_User" u LEFT JOIN "K_User_in_Project" up on u.k_user_id = up.user_id 
			WHERE up.project_id = ?`
	//return "SELECT user_id, user_name, head_shot_url FROM \"K_User\" WHERE user_id in " +
	//	"(SELECT user_id FROM \"K_User_in_Project\" WHERE project_id = ?)"
}

//查询全部管理的项目
//role_id >= 3拥有项目管理权限
func getAllManagedProjectsQuery() string {
	return "select * from \"K_Project\" where project_id in " +
		"(select project_id from \"K_User_in_Project\" where user_id = ? AND role_id >= 3)"
}

