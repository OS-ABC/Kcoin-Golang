package service

import (
	"Kcoin-Golang/src/models"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// struct for github api: "https://api.github.com/repos/"+userName+"/"+programName发送请求后的返回值
type JsonData struct {
	ID       int    `json:"id"`
	NodeID   string `json:"node_id"`
	Name     string `json:"name"`
	FullName string `json:"full_name"`
	Private  bool   `json:"private"`
	Owner    struct {
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
	} `json:"owner"`
	HTMLURL          string      `json:"html_url"`
	Description      interface{} `json:"description"`
	Fork             bool        `json:"fork"`
	URL              string      `json:"url"`
	ForksURL         string      `json:"forks_url"`
	KeysURL          string      `json:"keys_url"`
	CollaboratorsURL string      `json:"collaborators_url"`
	TeamsURL         string      `json:"teams_url"`
	HooksURL         string      `json:"hooks_url"`
	IssueEventsURL   string      `json:"issue_events_url"`
	EventsURL        string      `json:"events_url"`
	AssigneesURL     string      `json:"assignees_url"`
	BranchesURL      string      `json:"branches_url"`
	TagsURL          string      `json:"tags_url"`
	BlobsURL         string      `json:"blobs_url"`
	GitTagsURL       string      `json:"git_tags_url"`
	GitRefsURL       string      `json:"git_refs_url"`
	TreesURL         string      `json:"trees_url"`
	StatusesURL      string      `json:"statuses_url"`
	LanguagesURL     string      `json:"languages_url"`
	StargazersURL    string      `json:"stargazers_url"`
	ContributorsURL  string      `json:"contributors_url"`
	SubscribersURL   string      `json:"subscribers_url"`
	SubscriptionURL  string      `json:"subscription_url"`
	CommitsURL       string      `json:"commits_url"`
	GitCommitsURL    string      `json:"git_commits_url"`
	CommentsURL      string      `json:"comments_url"`
	IssueCommentURL  string      `json:"issue_comment_url"`
	ContentsURL      string      `json:"contents_url"`
	CompareURL       string      `json:"compare_url"`
	MergesURL        string      `json:"merges_url"`
	ArchiveURL       string      `json:"archive_url"`
	DownloadsURL     string      `json:"downloads_url"`
	IssuesURL        string      `json:"issues_url"`
	PullsURL         string      `json:"pulls_url"`
	MilestonesURL    string      `json:"milestones_url"`
	NotificationsURL string      `json:"notifications_url"`
	LabelsURL        string      `json:"labels_url"`
	ReleasesURL      string      `json:"releases_url"`
	DeploymentsURL   string      `json:"deployments_url"`
	CreatedAt        time.Time   `json:"created_at"`
	UpdatedAt        time.Time   `json:"updated_at"`
	PushedAt         time.Time   `json:"pushed_at"`
	GitURL           string      `json:"git_url"`
	SSHURL           string      `json:"ssh_url"`
	CloneURL         string      `json:"clone_url"`
	SvnURL           string      `json:"svn_url"`
	Homepage         interface{} `json:"homepage"`
	Size             int         `json:"size"`
	StargazersCount  int         `json:"stargazers_count"`
	WatchersCount    int         `json:"watchers_count"`
	Language         string      `json:"language"`
	HasIssues        bool        `json:"has_issues"`
	HasProjects      bool        `json:"has_projects"`
	HasDownloads     bool        `json:"has_downloads"`
	HasWiki          bool        `json:"has_wiki"`
	HasPages         bool        `json:"has_pages"`
	ForksCount       int         `json:"forks_count"`
	MirrorURL        interface{} `json:"mirror_url"`
	Archived         bool        `json:"archived"`
	Disabled         bool        `json:"disabled"`
	OpenIssuesCount  int         `json:"open_issues_count"`
	License          interface{} `json:"license"`
	Forks            int         `json:"forks"`
	OpenIssues       int         `json:"open_issues"`
	Watchers         int         `json:"watchers"`
	DefaultBranch    string      `json:"default_branch"`
	Organization     struct {
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
	} `json:"organization"`
	NetworkCount     int `json:"network_count"`
	SubscribersCount int `json:"subscribers_count"`
}

// struct for github api: "https://api.github.com/repos/" + userName + "/" + programName + "/" + "contributors"
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

// struct for github api: "https://api.github.com/users/" + username
type API_User struct {
	Login               string `json:"login"`
	Id                  int    `json:"id"`
	Node_id             string `json:"node_id"`
	Avatar_url          string `json:"avatar_url"`
	Gravatar_id         string `json:"gravatar_id"`
	Url                 string `json:"url"`
	Html_url            string `json:"html_url"`
	Followers_url       string `json:"followers_url"`
	Following_url       string `json:"following_url"`
	Gists_url           string `json:"gists_url"`
	Starred_url         string `json:"starred_url"`
	Subscriptions_url   string `json:"subscriptions_url"`
	Organizations_url   string `json:"organizations_url"`
	Repos_url           string `json:"repos_url"`
	Events_url          string `json:"events_url"`
	Received_events_url string `json:"received_events_url"`
	Type                string `json:"type"`
	Site_admin          bool   `json:"site_admin"`
	Name                string `json:"name"`
	Company             string `json:"company"`
	Blog                string `json:"blog"`
	Location            string `json:"location"`
	Email               string `json:"email"`
	Hireable            bool   `json:"hireable"`
	Bio                 string `json:"bio"`
	Public_repos        int    `json:"public_repos"`
	Public_gists        int    `json:"public_gists"`
	Followers           int    `json:"followers"`
	Following           int    `json:"following"`
	Created_at          string `json:"created_at"`
	Updated_at          string `json:"updated_at"`
}

/**
 * 这是一个全局数据结构，目前只有三个字段，用来保存GithubID, Github Name和对应access_token
 */
type GithubInfo struct {
	GithubId    string
	GithubName  string
	AccessToken string
}

// TODO 妥善使用该数据结构, 用户信息应该用session保存, 可以建立一个sessionID->session的映射, 但是不太清楚session如何使用, 这里需要会的人来修改这个丑陋的数据结构
//Github UserID -> GithubInfo
type GithubUserMap map[string]*GithubInfo

var GithubUser GithubUserMap

func init() {
	fmt.Println("Controller initialized!")
	GithubUser = make(GithubUserMap)
}

/**
 * 获取项目star数量的接口
 * 函数名：GetStarNum
 * 函数参数：url string
 * 返回值：starNum int 返回url对应项目的star数目
 */
func GetStarNum(url string) int {
	var starNum = 0
	//从url中获取到用户名和项目名
	userName, programName, _ := ParseGithubHTTPSUrl(url)

	var Api = "https://api.github.com/repos/" + userName + "/" + programName
	client := &http.Client{}
	response, _ := client.Get(Api)
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	var jd JsonData
	json.Unmarshal(body, &jd)
	starNum = jd.StargazersCount
	return starNum
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

type Data struct {
	Id   string `json:userId`
	Name string `json:"userName"`
	Uri  string `json:"headShotUrl"`
}

type UserJson struct {
	ErrorCode int
	Data      Data
}

func GetUserJson(access_token string) UserJson {
	var url_2 = "https://api.github.com/user?" + "access_token=" + access_token

	client_2 := &http.Client{}
	response_2, _ := client_2.Get(url_2)
	defer response_2.Body.Close()
	body_2, err_2 := ioutil.ReadAll(response_2.Body)
	if err_2 != nil {
		panic(err_2)
	}

	// 获取ID
	var name = strings.Split(strings.Split(string(body_2), ",")[0], "\"")[3]
	var uri = strings.Split(strings.Split(string(body_2), ",")[3], "\"")[3]
	var id = strings.Split(strings.Split(string(body_2), ",")[1], ":")[1]

	//select id according to name
	data := Data{
		Name: name,
		Uri:  uri,
		Id:   id,
	}
	userJson := UserJson{
		ErrorCode: 0,
		Data:      data,
	}

	return userJson
}

func GetAccessToken(code string) (accessToken string, err error) {
	client_id := beego.AppConfig.String("client_id")
	client_secret := beego.AppConfig.String("client_secret")
	url_1 := "https://github.com/login/oauth/access_token?code=" + code + "&client_id=" + client_id + "&client_secret=" + client_secret

	client := &http.Client{}
	response, err := client.Get(url_1)
	if err != nil {
		return "", err
	} else {
		defer response.Body.Close()
	}
	body, err1 := ioutil.ReadAll(response.Body)
	if err1 != nil {
		panic(err1)
	}
	accessToken = strings.Split(strings.Split(string(body), "&")[0], "=")[1]
	return accessToken, err
}

/**
 * 设置Github User这个map的Access Token字段.
 */
func (this GithubUserMap) SetGithubUserAccessToken(id string, name string, accessToken string) {
	if _, ok := this[id]; !ok {
		this[id] = new(GithubInfo)
	}
	this[id].AccessToken = accessToken
	this[id].GithubId = id
	this[id].GithubName = name
}

func (this GithubUserMap) GetGithubUserAccessToken(userId string) (string, error) {
	if userInfo, ok := this[userId]; ok {
		return userInfo.AccessToken, nil
	} else {
		return "", fmt.Errorf("user id %s is not valid", userId)
	}
}

//getWebhooksUrl 可以通过
func RegisterGithubWebhooks(userId string, repoName string) {
	accessToken, _ := GithubUser.GetGithubUserAccessToken(userId)
	postPayload := getPayloadOfRegisterGithubWebhooks()
	userName := GithubUser[userId].GithubName
	api_url := getWebhooksUrlBy(userName, repoName)
	bytePostPayload := []byte(postPayload)
	buffer := bytes.NewBuffer(bytePostPayload)
	request, err := http.NewRequest("POST", api_url, buffer)
	if err != nil {
		fmt.Printf("http.NewRequest%v", err)
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "token "+accessToken)
	client := &http.Client{}
	resp, err := client.Do(request.WithContext(context.TODO()))
	if err != nil {
		fmt.Printf("client.Do%v", err)
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		fmt.Printf("ioutil.ReadAll%v", err)
	}

	fmt.Println(string(respBytes))
}

/**
 * TODO 在config中设置secret, 与github_webhooks中的Post函数一起完成
 */
func getPayloadOfRegisterGithubWebhooks() string {
	return `{
  "name": "web",
  "active": true,
  "events": [
    "push",
    "pull_request"
  ],
  "config": {
    "url": "http://114.115.206.8:8080/webhooks",
    "content_type": "json",
    "insecure_ssl": "0"
  }
}`
}

func getWebhooksUrlBy(userName string, repoName string) string {
	return "https://api.github.com/repos/" + userName + "/" + repoName + "/hooks"
}

func GetGithubId(username string) int {
	api := "https://api.github.com/users/" + username
	client := &http.Client{}
	response, _ := client.Get(api)
	defer response.Body.Close()
	body, err_1 := ioutil.ReadAll(response.Body)
	if err_1 != nil {
		panic(err_1)
	}
	var res API_User
	json.Unmarshal(body, &res)
	fmt.Println(res.Id)
	return res.Id
}

type GithubRepos []struct {
	Name string `json:"name"`
	Url  string `json:"html_url"`
}

type GithubOrgUrl []struct{
	Url  string `json:"repos_url"`
}

type GithubOrgName []struct{
	Name  string `json:"login"`
}

/**
 * 解析github的reposUrl所对应的所有repos
 * 参数：reposUrl
 * 返回值：repos 包括项目名和url
 */
func GetPersonalRepos(reposUrl string) GithubRepos {
	var repos GithubRepos
	client := &http.Client{}
	// 根据reposUrl发送GET请求
	response, _ := client.Get(reposUrl)
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	// 解析json
	err1 := json.Unmarshal([]byte(body), &repos)
	if err1 != nil {
		panic(err1)
	}
	return repos
}

/**
 * 获取个人公开repos和所属org下的repos
 * 参数：user 用户名
 * 返回值：ProjectInfo json数据
 */
func GetGithubRepos(user string) (string, error) {
	var orgUrl string = "https://api.github.com/users/" + user + "/orgs"
	var reposUrl string = "https://api.github.com/users/" + user + "/repos"
	var orgs GithubOrgUrl
	var projects models.ProjectInfo
	client := &http.Client{}
	// 根据reposUrl发送GET请求
	response, _ := client.Get(orgUrl)
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	// 解析用户所属org
	err1 := json.Unmarshal([]byte(body), &orgs)
	if err1 != nil {
		panic(err1)
	}
	// 用户个人公开repos
	repos := GetPersonalRepos(reposUrl)
	// 从每个orgs中获取repos
	for _, value := range orgs{
		tmp := GetPersonalRepos(value.Url)
		for _, v := range tmp{
			repos = append(repos, v)
		}
	}

	for i := range repos {
		p := &models.Project{}
		p.ProjectUrl = repos[i].Url
		p.ProjectName = repos[i].Name
		projects.Data = append(projects.Data, p)
	}
	res, err2 := json.Marshal(&projects)
	if err2 != nil {
		panic(err2)
	}
	return string(res), nil
}

/**
 * 获取用户所属的所有组织name
 * 参数：user 用户名
 * 返回值：res 组织名字符串，使用空格拼接
 */
func GetOrgNames(user string) string {
	var url string = "https://api.github.com/users/" + user + "/orgs"
	var orgNames GithubOrgName
	res := ""
	client := &http.Client{}
	response, _ := client.Get(url)
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	err1 := json.Unmarshal([]byte(body), &orgNames)
	if err1 != nil {
		panic(err1)
	}
	for i, v := range orgNames{
		if i == len(orgNames)-1{
			res = res + v.Name
		} else {
			res = res + v.Name + " "
		}
	}
	return res
}