package models

import (
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

func init() {
	_ = orm.RegisterDriver("postgres", orm.DRPostgres)
	err := orm.RegisterDataBase("default", "postgres", "user=sspkukcoin password=kcoin2019 dbname=postgres host=114.115.133.140 port=5432")
	if err != nil {
		fmt.Println(err.Error())
	}
}

func GetUserInfo(userName string) *KUser {
	o := orm.NewOrm()
	_ = o.Using("default")
	result := new(KUser)
	err := o.Raw("SELECT * FROM \"K_User\" where user_name = ?", userName).QueryRow(result)
	if err != nil {
		fmt.Println(err.Error())
	}
	return result
}
