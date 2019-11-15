package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

func init() {
	_ = orm.RegisterDriver("postgres", orm.DRPostgres)
	err := orm.RegisterDataBase("default", "postgres", "user=sspkukcoin password=kcoin2019 dbname=postgres host=114.115.133.140 port=5432")
	if err != nil {
		fmt.Println(err.Error())
	}

	orm.RegisterModel(new(KProject))
}
