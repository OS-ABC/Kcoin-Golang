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

type projectInfo struct {
	ErrorCode string     `json:"errorCode"`
	Projects  []KProject `json:"data"`
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(KProject))
}

//查询并以json形式返回所有的项目信息
func GetAllProjectsInfo() (string, error) {
	o := orm.NewOrm()
	o.Using("default")

	var projectsInfo projectInfo
	projectsInfo.ErrorCode = "default Error"

	//var projects []KProject
	querySql := `SELECT * FROM "K_Project"`
	_, err := o.Raw(querySql).QueryRows(&projectsInfo.Projects)
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
