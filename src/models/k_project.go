package models

import (
	"encoding/json"
	_ "encoding/json"
	"fmt"
	"github.com/astaxie/beego/orm"
)

//查询并以json形式返回所有的项目信息
func GetAllProjectsInfo() (string, error) {
	o := orm.NewOrm()
	o.Using("default")

	var projectsInfo ProjectInfo
	//var projects []Project
	projectsInfo.ErrorCode = "default Error"

	//var projects []KProject
	querySql := `SELECT project_id , project_name , project_url , project_cover_url FROM "K_Project"`
	_, err := o.Raw(querySql).QueryRows(&projectsInfo.Data)
	for i := range projectsInfo.Data {
		u := &UserData{}
		u.HeadShotUrl = "../static/img/tx1.png"
		u.UserName = "abc"
		projectsInfo.Data[i].MemberList = append(projectsInfo.Data[i].MemberList, u)
	}
	if err != nil {
		fmt.Println(err.Error())
		return fmt.Sprint(projectsInfo), err
	}
	projectsInfo.ErrorCode = "0"
	res, err := json.Marshal(&projectsInfo)
	if err != nil {
		fmt.Println(err.Error())
		return fmt.Sprint(projectsInfo), err
	}
	return string(res), nil
}
