package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

//定义结构体，用来对接前后端json
//homepage
type Project struct {
	ProjectId string			`json:"projectId"`
	ProjectName string          `json:"projectName"`
	ProjectCoverUrl string      `json:"projectCoverUrl"`
	ProjectUrl string           `json:"projectUrl"`
	MemberList []*UserData       `json:"memberList"`
}

type ProjectInfo struct {
	ErrorCode string     `json:"errorCode"`
	Data  []*Project `json:"data"`
}

//ccSearchPage
/*UserCcData:为ccSearchPage的主要结构体，包括kcoin记录中的cc操作日期，操作类型和cc的数量
 */
 type UserCcOpe struct{
	OpeCcDate 	string	 `json:"opeCcDate"`
	OpeCcType	string	 `json:"opeCcType"`
	OpeCcNumber	string	 `json:"opeCcNumber"`
}

//homepage import
/*UserData:为personalpage和projectpage的主要结构体，定义了用户姓名、用户头像url、项目列表
 */
type UserData struct {
	UserId		string    `json:"userId"`
	UserName    string    `json:"userName"`
	HeadShotUrl string    `json:"headshotUrl"`
	ProjectList []*Project `json:"projectList"`
	UserCcOpeList []*UserCcOpe `json:"userCcOpeList"`
	CsNum       int       `json:"csNum"`
}

//import personPage
/*UserInfo:包括UserData和另外一个errorCode，errorCode主要用于调试
 */
type UserInfo struct {
	ErrorCode string   `json:"errorCode"`
	Data      *UserData `json:"data"`
}

func init() {
	_ = orm.RegisterDriver("postgres", orm.DRPostgres)
	err := orm.RegisterDataBase("default", "postgres", "user=sspkukcoin password=kcoin2019 dbname=postgres host=114.115.133.140 port=5432")
	if err != nil {
		fmt.Println(err.Error())
	}
}
