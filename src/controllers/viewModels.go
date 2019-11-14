package controllers

//定义结构体，用来对接前后端json

//homepage
type Project struct {
	ProjectName string          `json:"projectName"`
	ProjectCoverUrl string      `json:"projectCoverUrl"`
	ProjectUrl string           `json:"projectUrl"`
	MemberList []UserData       `json:"memberList"`
}
//homepage
//对应Json中数据结构的结构体
type ProjectList struct {
	ErrorCode string            `json:"errorCode"`
	Data []Project		        `json:"data"`
}

//homepage import
/*UserData:为personalpage和projectpage的主要结构体，定义了用户姓名、用户头像url、项目列表
 */
type UserData struct {
	UserName    string    `json:"userName"`
	HeadShotUrl string    `json:"headshotUrl"`
	ProjectList []Project `json:"projectList"`
}

//import personPage
/*UserInfo:包括UserData和另外一个errorCode，errorCode主要用于调试
 */
type UserInfo struct {
	ErrorCode string   `json:"errorCode"`
	Data      UserData `json:"data"`
}

