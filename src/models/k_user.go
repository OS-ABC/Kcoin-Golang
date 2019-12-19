package models

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
)

type noResultErr int

func (err noResultErr) Error() string {
	return "There is no result."
}

func GetUserInfo(userName string) (string, error) {
	userInfo := UserInfo{}
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
			userData.HeadShotUrl = maps[i]["head_shot_url"].(string)
		}
	}
	userInfo.Data = userData

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
SELECT DISTINCT(project_id) AS attended_pro_id from "k_user_in_project"
where user_id in(
SELECT user_id from "k_user"
where user_name=?
)
)

SELECT a.project_name,a.project_url,a.project_cover_url,b.user_name,b.head_shot_url FROM

(SELECT project_id, project_name,project_url,project_cover_url FROM "k_project" kpro
where kpro.project_id in
(SELECT attended_pro_id FROM attend_project)
) a

LEFT JOIN

(SELECT kuip.project_id,ku.user_name,ku.head_shot_url FROM "k_user_in_project" kuip INNER JOIN "k_user" ku ON kuip.user_id=ku.user_id
where kuip.project_id in
(SELECT attended_pro_id FROM attend_project)
) b

ON a.project_id=b.project_id`
}

func buildProjectsDataFrom(maps []orm.Params) *UserData {
	userData := &UserData{}
	projectToIndex := make(map[string]int)
	index := 0
	for i := range maps {
		projectName := maps[i]["project_name"].(string)
		if _, ok := projectToIndex[projectName]; !ok {
			// 如果project name不存在, 那么添加这个projectName, 并且新建这个Project的信息
			projectToIndex[projectName] = index
			index++
			projectInfo := &Project{}
			projectInfo.ProjectName = projectName
			projectInfo.ProjectUrl = maps[i]["project_url"].(string)
			projectInfo.ProjectCoverUrl = maps[i]["project_cover_url"].(string)
			userData.ProjectList = append(userData.ProjectList, projectInfo)
		}
		member := &UserData{}
		member.UserName = maps[i]["user_name"].(string)
		member.HeadShotUrl = maps[i]["head_shot_url"].(string)
		projectIndex := projectToIndex[projectName]
		userData.ProjectList[projectIndex].MemberList = append(userData.ProjectList[projectIndex].MemberList, member)
	}
	return userData
}

func jsonize(info UserInfo) (string, error) {
	if res, err := json.Marshal(&info); err != nil {
		fmt.Println(err.Error())
		return fmt.Sprint(info), err
	} else {
		return string(res), nil
	}
}

func FinduserByGitId(id string) (UserData, error) {
	o := orm.NewOrm()
	_ = o.Using("default")
	querySql := `select * from "k_user" where github_user_id = ?`
	var maps []orm.Params
	var u = UserData{}
	_, err := o.Raw(querySql, id).Values(&maps)
	if err != nil {
		fmt.Println(err.Error())
	}
	for i := range maps {
		u.UserId = maps[i]["k_user_id"].(string)
		u.UserName = maps[i]["user_name"].(string)
	}
	return u, err
}
func FindUserByUsername(username string) (sql.Result, error) {
	o := orm.NewOrm()
	_ = o.Using("default")
	querySql := `select k_user_id from "k_user" where user_name = ?`
	res, err := o.Raw(querySql, username).Exec()
	return res, err
}
func GetUseridByUsername(username string) (int, error) {
	var k_user_id int
	o := orm.NewOrm()
	_ = o.Using("default")
	querySql := `select k_user_id from "k_user" where user_name = ?`
	err := o.Raw(querySql, username).QueryRow(&k_user_id)
	return k_user_id, err
}
func InsertIntoKUserInProject(projectId int, userId int) (sql.Result, error) {
	o := orm.NewOrm()
	_ = o.Using("default")
	insertSql := `insert into "k_user_in_project" (project_id,user_id)values(?,?)`
	res, err := o.Raw(insertSql, projectId, userId).Exec()
	return res, err
}
func InsertIntoKTemporaryUser(inviterId, gitId int, gitName string, projectId int) (sql.Result, error) {
	o := orm.NewOrm()
	_ = o.Using("default")
	insertSql := `insert into "k_temporary_user" (inviter_id,git_id,git_name,project_id,invite_time)values(?,?,?,?,?)`
	currentTime := time.Now()
	currentTime.Format("2006-01-02 15:04:05:000000")
	res, err := o.Raw(insertSql, inviterId, gitId, gitName, projectId, currentTime).Exec()
	return res, err
}
func InsertUser(name string, uri string, id string) error {
	o := orm.NewOrm()
	_ = o.Using("default")
	time := time.Now().Format("2006-01-02 15:04:05.000000")

	insertSql := `INSERT INTO "k_user" (USER_NAME,REGISTER_TIME,HEAD_SHOT_URL,GITHUB_USER_ID) VALUES (?,?,?,?);`
	_, err := o.Raw(insertSql, name, time, uri, id).Exec()

	return err
}
func IsSupervisor(id string) bool {
	o := orm.NewOrm()
	_ = o.Using("default")
	findSql := `select * from "k_supervisor" where k_user_id= ?`
	res, _ := o.Raw(findSql, id).Exec()
	if res == nil {
		return false
	} else if n, _ := res.RowsAffected(); n == 0 {
		return false
	}
	return true
}
func FindUserInKUserInProject(userid int) (int, error) {
	o := orm.NewOrm()
	_ = o.Using("default")
	querySql := `select project_id from "k_user_in_project" where user_id = ?`
	res, err := o.Raw(querySql, userid).Exec()
	num, err := res.RowsAffected()
	return int(num), err
}

func InsertKCsChangeRecord(projectId int,projectName string,acceptUserId int,acceptUserName string,csAmount float64)(sql.Result,error){
	o:=orm.NewOrm()
	insertSql:=`insert into "k_cs_change_record"(distribute_project_id,distribute_project_name,accept_user_id,accept_user_name,cs_amount,distribute_time)values(?,?,?,?,?,?)`
	currentTime := time.Now()
	currentTime.Format("2006-01-02 15:04:05:000000")
	res, err := o.Raw(insertSql, projectId, projectName, acceptUserId, acceptUserName, csAmount,currentTime).Exec()
	return  res,err

}


//该函数通过项目id获取该项目的所有成员信息
func GetMembersInfoByProjectName(projectName string) (membersInfo []*UserData, err error) {
	var memberlist []*UserData
	var projectid int
	o := orm.NewOrm()
	o.Using("default")
	queryProjectIDSql := getProjectIDQuery()
	if err = o.Raw(queryProjectIDSql, projectName).QueryRow(&projectid); err != nil {
		fmt.Print(err.Error())
		return nil, err
	}

	queryMembersInProjectSql := getAllMemberQuery()
	if _, err = o.Raw(queryMembersInProjectSql, projectid).QueryRows(&memberlist); err != nil {
		fmt.Print(err.Error())
		return nil, err
	}

	return memberlist, nil
}

//通过连接k_user表和k_user_in_project表查询用户信息
func getAllMemberQuery() string {
	return `SELECT u.k_user_id, u.user_name, u.head_shot_url
			FROM "k_user" u LEFT JOIN "k_user_in_project" up on u.k_user_id = up.user_id 
			WHERE up.project_id = ?`
}

//ProjectId只能通过查询K_project表获取，所以getProjectId函数通过函数名查询ProjectId后返回
func getProjectIDQuery() string {
	return `SELECT project_id FROM "k_project" WHERE project_name = ?`
}

