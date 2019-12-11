package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)
type CS struct {
	sort string
	content string
	value int
}

/*
获取通过提交pr得到的CS数和具体的pr信息
 */
func getCS_pr(project string,member string,x int) []CS{
	var slice=[]CS{}

	var cs_pr CS
	cs_pr.sort="pr"
	cs_pr.content="提交测试pr"
	cs_pr.value=x

	slice=append(slice,cs_pr)
	return slice
}

/*
获取通过合并pr得到的CS数和具体的pr信息
*/
func getCS_mergePr(project string,member string,x int) []CS{
	var slice=[]CS{}

	var cs_mergePr CS
	cs_mergePr.sort="合并pr"
	cs_mergePr.content="合并测试pr"
	cs_mergePr.value=x

	slice=append(slice,cs_mergePr)
	return slice
}

/*
获取通过review他人pr得到的CS数和具体的pr信息
*/
func getCS_review(project string,member string,y int) []CS{
	var slice=[]CS{}

	var cs_review CS
	cs_review.sort="review Pr"
	cs_review.content="review测试pr"
	cs_review.value=y
	slice=append(slice,cs_review)

	return slice
}

/*
获取通过提交commit信息得到的CS数和具体的commit信息
*/
func getCS_commit(project string,member string,z int) []CS{
	var slice=[]CS{}

	var cs_commit CS
	cs_commit.sort="commit"
	cs_commit.content="测试commit"
	cs_commit.value=z
	slice=append(slice,cs_commit)

	return slice
}

/*
获取通过提交issue得到的CS数和具体的issue信息
*/
func getCS_issue(project string,member string,m int) []CS{
	var slice=[]CS{}

	var cs_issue CS
	cs_issue.sort="commit"
	cs_issue.content="测试commit1"
	cs_issue.value=m
	slice=append(slice,cs_issue)

	return slice
}

/*
获取CS数和CS贡献度列表
*/
func getCS(project string,member string,x int,y int,z int,m int) []CS{
	var cs_total=[]CS{}
	cs_total = append(getCS_pr(project, member, x), getCS_mergePr(project, member, x)...)
	cs_total = append(cs_total, getCS_review(project,member,y)...)
	cs_total = append(cs_total, getCS_commit(project,member,z)...)
	cs_total = append(cs_total, getCS_issue(project,member,m)...)
	return cs_total
}

//func main() {
//	var info =getCS("OS-ABC","Kcoin-Golang",1,2,3,1)
//	fmt.Println(info)
//}


