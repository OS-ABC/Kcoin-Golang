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
	queryProjectSql := `SELECT project_id , project_name , project_url , project_cover_url FROM "K_Project" `
	/***********************************************************************************************************/

	_, err := o.Raw(queryProjectSql).QueryRows(&projectsInfo.Data)

	/******************************************query menberList in one project**********************************/
	queryUsersInProjectSql := `select u.k_user_id,u.user_name,u.head_shot_url
								from "K_User" u left join "K_User_in_Project" up on u.k_user_id=up.user_id
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

	queryProjectSql := `SELECT project_cc FROM "K_Project" WHERE project_name=? `
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
	querySql := `select project_id from "K_Project" where project_name=?`
	err := o.Raw(querySql, reponame).QueryRow(&project_id)
	return project_id, err
}

func InsertProject(reponame string, url string, project_cover_url string) (sql.Result, error) {
	o := orm.NewOrm()
	_ = o.Using("default")
	querySql := `insert into "K_Project"(project_name,project_url,project_cover_url)values(?,?,?)`
	res, err := o.Raw(querySql, reponame, url, project_cover_url).Exec()
	if err != nil {
		log.Fatal("error when insert project,", err)
	}
	return res, err
}

/*
查询托管在平台的项目总数
不需要传入参数，返回查询的结果（项目总数）和错误信息
*/
func PlatfmProjNum() (int64, error) {
	o := orm.NewOrm()
	_ = o.Using("default")
	// return o.QueryTable("K_Project").Count()
	type Row struct{
		Project_id int
	}
	var rows []Row
	projNumQuery := `SELECT project_id FROM "K_Project"`

	// 这句将结果赋值给rows, 并且将查到的行数（即项目总数）赋值给num
	// 如果有错误，赋值给err。项目包含KCoin项目本身
	// 用高级查询里的Count()本来可以不浪费空间存储结构体，可是报错说表名不存在？
	num, err := o.Raw(projNumQuery).QueryRows(&rows)

	if err != nil {
		fmt.Println("Error occured:", err.Error())
		return -1, err
	}
	return num, nil
}