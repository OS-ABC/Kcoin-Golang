package models
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//解析api获取到的信息，需要的用户名是其中的"login"
type ContributorData []struct {
	Login string `json:"login"`
	ID int `json:"id"`
	NodeID string `json:"node_id"`
	AvatarURL string `json:"avatar_url"`
	GravatarID string `json:"gravatar_id"`
	URL string `json:"url"`
	HTMLURL string `json:"html_url"`
	FollowersURL string `json:"followers_url"`
	FollowingURL string `json:"following_url"`
	GistsURL string `json:"gists_url"`
	StarredURL string `json:"starred_url"`
	SubscriptionsURL string `json:"subscriptions_url"`
	OrganizationsURL string `json:"organizations_url"`
	ReposURL string `json:"repos_url"`
	EventsURL string `json:"events_url"`
	ReceivedEventsURL string `json:"received_events_url"`
	Type string `json:"type"`
	SiteAdmin bool `json:"site_admin"`
	Contributions int `json:"contributions"`
}


//获取项目贡献者信息的接口，返回string,不同contributor在字符串中用" "分割
func getContributors(userName string,programName string) string {
	var url_1 string = "https://api.github.com/repos/"+userName+"/"+programName+"/"+"contributors"

	client :=&http.Client{}
	response,_:=client.Get(url_1)
	defer response.Body.Close()
	body,err_1:=ioutil.ReadAll(response.Body)
	if err_1 != nil{
		panic(err_1)}

	var cb ContributorData
	json.Unmarshal(body, &cb)
	var cl string=""
	for i:= 0;i<len(cb);i++{
		var Name string=cb[i].Login
		cl=cl+Name+" "
	}
	fmt.Println(cl)
	return cl
}

//获取项目贡献者的人数
func GetContributorNum(url string) int {

	str1 := strings.Split(url, "https://github.com/")[1]
	str2 := strings.Split(str1, "/")
	user_Name := str2[0]
	program_Name := str2[1]

	info := getContributors(user_Name,program_Name)
	res := strings.TrimSpace(info)
	str_arr := strings.Split(res, " ")
	count := len(str_arr)
	return count
}
