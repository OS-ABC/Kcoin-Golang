package models

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

type KUser struct {
	Id           int64 `orm:"pk;column(user_id);"`
	UserName     string
	RegisterTime int
	HeadShotUrl  string
}

type userInfo struct {
	ErrorCode string   `json:"errorCode"`
	UserData  userData `json:"data"`
}

type userBase struct {
	UserName    string `json:"userName"`
	HeadShotURL string `json:"headShotUrl"`
}

type userData struct {
	userBase
	Projects []project `json:"projectList"`
}

type project struct {
	ProjectName     string     `json:"projectName"`
	ProjectCoverUrl string     `json:"projectCoverUrl"`
	ProjectUrl      string     `json:"projectUrl"`
	Members         []userBase `json:"memberList"`
}

/**
* select * from "K_User"
* left join "K_User_in_Project"
* on "K_User".user_id = "K_User_in_Project".user_id;
 */

func init() {
	_ = orm.RegisterDriver("postgres", orm.DRPostgres)
	err := orm.RegisterDataBase("default", "postgres", "user=sspkukcoin password=kcoin2019 dbname=postgres host=114.115.133.140 port=5432")
	if err != nil {
		fmt.Println(err.Error())
	}
}

func GetUserInfo(userName string) (string, error) {

	var userInfo userInfo
	userInfo.ErrorCode = "default Error"

	o := orm.NewOrm()
	_ = o.Using("default")
	var userData userData
	// TODO 修改SQL Query, 直接查询到userData, 主要有userName, headShotUrl和projectList,projectList中有project的详细信息
	/*****************************************************/
	SQLQuery := `SELECT * FROM "K_User" where user_name = ?`
	/*****************************************************/
	if err := o.Raw(SQLQuery, userName).QueryRow(userData); err != nil {
		fmt.Println(err.Error())
		return fmt.Sprint(userInfo), err
	}
	// Error Code may be ugly! Use err to Debug!
	userInfo.ErrorCode = "0"
	userInfo.UserData = userData
	if res, err := json.Marshal(&userInfo); err != nil {
		fmt.Println(err.Error())
		return fmt.Sprint(userInfo), err
	} else {
		return fmt.Sprint(res), nil
	}
}
