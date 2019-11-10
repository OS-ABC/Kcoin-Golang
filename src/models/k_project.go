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

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(KProject))
}

//查询并以json形式返回所有的项目信息
func GetAllProjectsInfo() string {
	o := orm.NewOrm()
	o.Using("default")

	var projects []KProject
	_, err := o.Raw("SELECT * FROM \"K_Project\"").QueryRows(&projects)
	if err != nil {
		fmt.Println(err)
	}
	res, err := json.Marshal(&projects)
	if err != nil {
		fmt.Println(err)
	}
	return string(res)
}
