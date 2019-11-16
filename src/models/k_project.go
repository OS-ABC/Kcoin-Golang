package models

import (
	"encoding/json"
	_ "encoding/json"
	"fmt"
	"github.com/astaxie/beego/orm"
)

type KProject struct {
	ProjectId       int64  `json:"pid" orm:"pk;column(project_id);"`
	ProjectName     string `json:"pname"`
	ProjectUrl      string `json:"purl"`
	ProjectCoverUrl string `json:"pcurl"`
}

type Project struct {
	ProjectName string          `json:"projectName"`
	ProjectCoverUrl string      `json:"projectCoverUrl"`
	ProjectUrl string           `json:"projectUrl"`
	MemberList []UserData       `json:"memberList"`
}
type UserData struct {
	UserName    string    `json:"userName"`
	HeadShotUrl string    `json:"headshotUrl"`
	ProjectList []Project `json:"projectList"`
}
type projectInfo struct {
	ErrorCode string     `json:"errorCode"`
	Projects  []Project `json:"data"`
}
type kprojectInfo struct {
	ErrorCode string     `json:"errorCode"`
	Projects  []KProject `json:"data"`
}

//查询并以json形式返回所有的项目信息
func GetAllProjectsInfo() (string, error) {
	o := orm.NewOrm()
	o.Using("default")

	var projectsInfo projectInfo
	//var projects []Project
	projectsInfo.ErrorCode = "default Error"

	//var projects []KProject
	querySql := `SELECT project_name , project_url , project_cover_url FROM "K_Project"`
	_, err := o.Raw(querySql).QueryRows(&projectsInfo.Projects)
	for i := range projectsInfo.Projects {
		var u UserData
		u.HeadShotUrl = "../static/img/tx1.png"
		u.UserName = "abc"
		projectsInfo.Projects[i].MemberList = append(projectsInfo.Projects[i].MemberList, u)
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
