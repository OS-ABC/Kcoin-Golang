package models

import (
	"database/sql"
	"encoding/json"
	_ "encoding/json"
	"fmt"
	"log"

	"github.com/astaxie/beego/orm"
)

//查询并以json形式返回所有的项目信息
func GetAllProjectsInfo() (string, error) {
	o := orm.NewOrm()
	o.Using("default")

	var projectsInfo ProjectInfo
	projectsInfo.ErrorCode = "default Error"

	/******************************************query all projects************************************************/
	queryProjectSql := `SELECT project_id , project_name , project_url , project_cover_url FROM "k_project" `
	/***********************************************************************************************************/

	_, err := o.Raw(queryProjectSql).QueryRows(&projectsInfo.Data)

	/******************************************query menberList in one project**********************************/
	queryUsersInProjectSql := `select u.k_user_id,u.user_name,u.head_shot_url
								from "k_user" u left join "k_user_in_project" up on u.k_user_id=up.user_id
       							where up.project_id=?`
	/**********************************************************************************************************/
	for _, v := range projectsInfo.Data {
		var memberList []*UserData
		_, err := o.Raw(queryUsersInProjectSql, v.ProjectId).QueryRows(&memberList)
		if err != nil {
			fmt.Println(err.Error())
			return fmt.Sprint(projectsInfo), err
		}
		v.MemberList = memberList
	}
	projectsInfo.ErrorCode = "0"
	res, err := json.Marshal(&projectsInfo)
	if err != nil {
		fmt.Println(err.Error())
		return fmt.Sprint(projectsInfo), err
	}
	return string(res), nil
}

func GetProjectsCC(projectName string) (float64, error) {
	o := orm.NewOrm()
	o.Using("default")

	queryProjectSql := `SELECT project_cc FROM "k_project" WHERE project_name=? `
	var num float64
	err := o.Raw(queryProjectSql, projectName).QueryRow(&num)
	if err != nil {
		fmt.Println(err.Error())
	}
	return num, nil
}

func GetProjectidByRepoName(reponame string) (int, error) {
	o := orm.NewOrm()
	_ = o.Using("default")
	var project_id int
	querySql := `select project_id from "k_project" where project_name=?`
	err := o.Raw(querySql, reponame).QueryRow(&project_id)
	return project_id, err
}

func InsertProject(reponame string, url string, project_cover_url string) (sql.Result, error) {
	o := orm.NewOrm()
	_ = o.Using("default")
	querySql := `insert into "k_project"(project_name,project_url,project_cover_url)values(?,?,?)`
	res, err := o.Raw(querySql, reponame, url, project_cover_url).Exec()
	if err != nil {
		log.Fatal("error when insert project,", err)
	}
	return res, err
}
