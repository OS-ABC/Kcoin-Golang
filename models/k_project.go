package models

import(
	"fmt"
)

// 导入项目时的信息
type ProjectDetail struct {
	ProjectName string     `gorm:"column:project_name" json:"projectName"`
	ProjectCoverUrl string `gorm:"column:project_cover_url" json:"projectCoverUrl"`
	Introduction string    `gorm:"column:project_description" json:"introdection"`
	GithubUrl string       `gorm:"column:project_url" json:"githubUrl"`
	DecideType int         `json:"decideType"` 
}

func AddProject(project *ProjectDetail) int {
	tem := -1
	//查询数据库是否已有该url的项目，若有则将其id保存在tem里
	DB.Table("k_project").Select("project_id").Where("project_url = ?", project.GithubUrl).First(&tem)
	if tem >= 0 { // 说明有ID，即项目已在平台内
		return 0
	} else if err := DB.Table("k_project").Create(&project).Error; err != nil {
		fmt.Println(err.Error())
		return -1
	} else {
		return 1
	}
}