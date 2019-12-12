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
func GetProjectidByRepoName(reponame string)(int,error){
	o:=orm.NewOrm()
	_ = o.Using("default")
	var project_id int
	querySql:=`select project_id from "K_Project" where project_name=?`
	err := o.Raw(querySql,reponame).QueryRow(&project_id)
	return project_id,err
}
func InsertProject(reponame string,url string,project_cover_url string)(sql.Result,error){
	o:=orm.NewOrm()
	_ = o.Using("default")
	querySql:=`insert into "K_Project"(project_name,project_url,project_cover_url)values(?,?,?)`
	res,err := o.Raw(querySql,reponame,url,project_cover_url).Exec()
	if err!=nil{
		log.Fatal("error when insert project,",err)
	}
	return res,err
}