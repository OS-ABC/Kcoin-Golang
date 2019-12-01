package models
import(
	"github.com/astaxie/beego/orm"
)
func SetProjectsDescription(projecturl string,description string) (errorcode error) {
	updatesql := `UPDATE K_Project SET project_url = "projecturl",project_description = "description" WHERE id = 4 ;`
	o := orm.NewOrm()
	o.Using("default")
	_,err := o.Raw(updatesql,projecturl , description).Exec()
	return err
}
