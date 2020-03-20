package models

import (
	"Kcoin-Golang/conf"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

// 用户信息结构体，tag里的json字段是与github返回的json名称保持一致，便于从github传回的json中把信息解析到结构体
type User struct {
	ID          int       `gorm:"primary_key;column:k_user_id" json:"kcoinID"`
	Name        string    `gorm:"column:user_name" json:"login"`
	Time        time.Time `gorm:"column:register_time" json:"registerTime"`
	HeadShotUrl string    `gorm:"column:head_shot_url" json:"avatar_url"`
	GithubID    int64     `gorm:"column:github_id" json:"id"`
}

//用户信息结构体，tag里的json字段是与需要返回给前端的json名称保持一致
type UserData struct {
	UserId      string `gorm:"primary_key;column:k_user_id" json:"userId"`
	UserName    string `json:"userName"`
	HeadShotUrl string `json:"headshotUrl"`
}

//项目集合结构体，用于封装获取到的全部项目集合，并作为json数据作为get请求的返回
type ProjectInfo struct {
	Projects []*Project `json:"projects"`
}

//项目信息结构体，tag里的json字段与需要返回给前端的json名称保持一致
type Project struct {
	ProjectID       int         `gorm:"primary_key" json:"projectId"`
	ProjectName     string      `json:"projectName"`
	ProjectCoverUrl string      `json:"projectCoverUrl"`
	MemberList      []*UserData `json:"memberList"`
}

func init() {
	var err error
	DB, err = gorm.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s", conf.Config.MySQL.Host, conf.Config.MySQL.Port, conf.Config.MySQL.User, conf.Config.MySQL.DBname, conf.Config.MySQL.Password))

	if err != nil {
		fmt.Println("open database failed, err: ", err)
	}
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "k_" + defaultTableName
	}
	// 禁用复数表名
	DB.SingularTable(true)
}
