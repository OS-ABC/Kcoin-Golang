package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

func CCInject(donatorId int, projId int, cc float64, dtype int) error {
	o := orm.NewOrm()
	err := o.Using("default")
	if err != nil {
		return fmt.Errorf("database connection error:"+err.Error())
	}
	SQLQuery := `INSERT INTO "k_cc_donate_record" (donate_from, donate_into, donate_cc, donate_type, donate_time)
				VALUES (?, ?, ?, ?, "CURRENT_TIMESTAMP(0)")`
	_, err = o.Raw(SQLQuery, donatorId, projId, cc, dtype).Exec()
	if err != nil {
		return fmt.Errorf("database insert error: "+err.Error())
	}
	err = nil
	return err
}
