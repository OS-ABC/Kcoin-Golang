package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

//解析api获取到的信息，需要的是labels中的name
type IssueData struct {
	URL           string `json:"url"`
	RepositoryURL string `json:"repository_url"`
	LabelsURL     string `json:"labels_url"`
	CommentsURL   string `json:"comments_url"`
	EventsURL     string `json:"events_url"`
	HTMLURL       string `json:"html_url"`
	ID            int    `json:"id"`
	NodeID        string `json:"node_id"`
	Number        int    `json:"number"`
	Title         string `json:"title"`
	User          struct {
		Login             string `json:"login"`
		ID                int    `json:"id"`
		NodeID            string `json:"node_id"`
		AvatarURL         string `json:"avatar_url"`
		GravatarID        string `json:"gravatar_id"`
		URL               string `json:"url"`
		HTMLURL           string `json:"html_url"`
		FollowersURL      string `json:"followers_url"`
		FollowingURL      string `json:"following_url"`
		GistsURL          string `json:"gists_url"`
		StarredURL        string `json:"starred_url"`
		SubscriptionsURL  string `json:"subscriptions_url"`
		OrganizationsURL  string `json:"organizations_url"`
		ReposURL          string `json:"repos_url"`
		EventsURL         string `json:"events_url"`
		ReceivedEventsURL string `json:"received_events_url"`
		Type              string `json:"type"`
		SiteAdmin         bool   `json:"site_admin"`
	} `json:"user"`
	Labels []struct {
		ID          int    `json:"id"`
		NodeID      string `json:"node_id"`
		URL         string `json:"url"`
		Name        string `json:"name"`
		Color       string `json:"color"`
		Default     bool   `json:"default"`
		Description string `json:"description"`
	} `json:"labels"`
	State             string        `json:"state"`
	Locked            bool          `json:"locked"`
	Assignee          interface{}   `json:"assignee"`
	Assignees         []interface{} `json:"assignees"`
	Milestone         interface{}   `json:"milestone"`
	Comments          int           `json:"comments"`
	CreatedAt         time.Time     `json:"created_at"`
	UpdatedAt         time.Time     `json:"updated_at"`
	ClosedAt          interface{}   `json:"closed_at"`
	AuthorAssociation string        `json:"author_association"`
	Body              string        `json:"body"`
	ClosedBy          interface{}   `json:"closed_by"`
}

//获取项目贡献者信息的接口
// 函数名：Get_issue_info
// 函数参数：userName string, programName string,issueName int
// 返回值：string 包含所有的labels信息

type Issue_result struct {
	issuer_id int
	issuer string
	csNum int
}

func Get_issue_info(userName string, programName string, issueNum int) (int, string, int) {
	Num := strconv.Itoa(issueNum)
	var url_1 string = "https://api.github.com/repos/" + userName + "/" + programName + "/" + "issues" + "/" + Num
	fmt.Println(url_1)
	client := &http.Client{}
	response, _ := client.Get(url_1)
	defer response.Body.Close()
	body, err_1 := ioutil.ReadAll(response.Body)
	if err_1 != nil {
		panic(err_1)
	}

	var ib IssueData
	json.Unmarshal(body, &ib)
	issues := []string{}
	for i := 0; i < len(ib.Labels); i++ {
		var labelsName string = ib.Labels[i].Name
		kcoin := "Kcoin"
		if strings.HasPrefix(labelsName, kcoin) == true {
			issues = append(issues, strings.Split(labelsName, "Kcoin#")[1])
		}
	}
	
	var res Issue_result
	res.issuer_id=ib.User.ID
	res.issuer=ib.User.Login
	res.csNum,_= strconv.Atoi(issues[0])
	return res.issuer_id, res.issuer, res.csNum
}

//使用实例：
//func main() {
//	var info int= Get_issue_info("OS-ABC", "Kcoin-Golang", 288)
//	fmt.Println(info)
//}
