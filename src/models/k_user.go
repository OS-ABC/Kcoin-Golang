package models

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/orm"
)

type noResultErr int

func (err noResultErr) Error() string {
	return "There is no result."
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

func GetUserInfo(userName string) (string, error) {
	var userInfo userInfo
	userInfo.ErrorCode = "default Error"

	maps, err := getSQLQueryResult(userName)
	if err != nil {
		return fmt.Sprint(userInfo), err
	}

	userData := buildProjectsDataFrom(maps)

	// complete userInfo building
	userInfo.ErrorCode = "0"
	userData.UserName = userName
	for i := range maps {
		if name := maps[i]["user_name"].(string); name == userName {
			userData.HeadShotURL = maps[i]["head_shot_url"].(string)
		}
	}
	userInfo.UserData = *userData

	return jsonize(userInfo)
}

func getSQLQueryResult(userName string) ([]orm.Params, error) {
	o := orm.NewOrm()
	_ = o.Using("default")
	SQLQuery := getUserInfoSQLQuery()
	var maps []orm.Params

	if _, err := o.Raw(SQLQuery, userName).Values(&maps); err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	if len(maps) == 0 {
		return nil, new(noResultErr)
	}
	return maps, nil
}

func getUserInfoSQLQuery() string {
	return `WITH attend_project AS
(
SELECT DISTINCT(project_id) AS attended_pro_id from "K_User_in_Project"
where user_id in(
SELECT User_id from "K_User"
where user_name=?
)
)

SELECT a.project_name,a.project_url,a.project_cover_url,b.user_name,b.head_shot_url FROM

(SELECT project_id, project_name,project_url,project_cover_url FROM "K_Project" kpro
where kpro.project_id in
(SELECT attended_pro_id FROM attend_project)
) a

LEFT JOIN

(SELECT kuip.project_id,ku.user_name,ku.head_shot_url FROM "K_User_in_Project" kuip INNER JOIN "K_User" ku ON kuip.user_id=ku.user_id
where kuip.project_id in
(SELECT attended_pro_id FROM attend_project)
) b

ON a.project_id=b.project_id`
}

func buildProjectsDataFrom(maps []orm.Params) *userData {
	userData := new(userData)
	projectToIndex := make(map[string]int)
	index := 0
	for i := range maps {
		projectName := maps[i]["project_name"].(string)
		if _, ok := projectToIndex[projectName]; !ok {
			// 如果project name不存在, 那么添加这个projectName, 并且新建这个Project的信息
			projectToIndex[projectName] = index
			index++
			var projectInfo project
			projectInfo.ProjectName = projectName
			projectInfo.ProjectUrl = maps[i]["project_url"].(string)
			projectInfo.ProjectCoverUrl = maps[i]["project_cover_url"].(string)
			userData.Projects = append(userData.Projects, projectInfo)
		}
		var member userBase
		member.UserName = maps[i]["user_name"].(string)
		member.HeadShotURL = maps[i]["head_shot_url"].(string)
		projectIndex := projectToIndex[projectName]
		userData.Projects[projectIndex].Members = append(userData.Projects[projectIndex].Members, member)
	}
	return userData
}

func jsonize(info userInfo) (string, error) {
	if res, err := json.Marshal(&info); err != nil {
		fmt.Println(err.Error())
		return fmt.Sprint(info), err
	} else {
		return string(res), nil
	}
}
