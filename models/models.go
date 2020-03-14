package models

import (
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

func init() {

	// TODO 修改从config文件中读取这些信息。
	var err error
	DB, err = gorm.Open("postgres", "host=114.115.133.140 port=5432 user=sspkukcoin dbname=postgres password=kcoin2019")
	if err != nil {
		fmt.Println("open database failed, err: ", err)
	}
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "k_" + defaultTableName
	}
	// 禁用复数表名
	DB.SingularTable(true)
}
