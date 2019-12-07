package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

//解析api获取到的信息，需要的用户名是其中的"login"
type ContributorData []struct {
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
	Contributions     int    `json:"contributions"`
}


//获取项目贡献者信息的接口
// 函数名：getContributors
// 函数参数：userName string programName string
// 返回值：string 包含所有的contributor信息，不同的contributor用" "分割

func GetContributors(userName string, programName string) string {
	var url_1 string = "https://api.github.com/repos/" + userName + "/" + programName + "/" + "contributors"

	client := &http.Client{}
	response, _ := client.Get(url_1)
	defer response.Body.Close()
	body, err_1 := ioutil.ReadAll(response.Body)
	if err_1 != nil {
		panic(err_1)
	}

	var cb ContributorData
	json.Unmarshal(body, &cb)
	var cl string = ""
	for i := 0; i < len(cb); i++ {
		var Name string = cb[i].Login
		cl = cl + Name + " "
	}
	fmt.Println(cl)
	return cl
}

//获取贡献者的人数
// 函数名：GetContributorNum
// 函数参数：url string
// 返回值：int 返回对应项目的贡献者人数
func GetContributorNum(url string) int {

	user_Name, program_Name, _ := ParseGithubHTTPSUrl(url)

	info := GetContributors(user_Name, program_Name)
	res := strings.TrimSpace(info)
	str_arr := strings.Split(res, " ")
	count := len(str_arr)
	return count
}