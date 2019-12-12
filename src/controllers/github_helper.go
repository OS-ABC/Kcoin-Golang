package controllers

import (
	"Kcoin-Golang/src/models"
	"bytes"
	"context"
	"encoding/json"
	_ "encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
	"net/http"
	"strings"
)

/**
 * 这是一个全局数据结构，目前只有两个字段，用来保存Github名和对应access_token
 */
type GithubInfo struct {
	GithubId    string
	GithubName  string
	AccessToken string
}
type API_User struct {
	Login string `json:"login"`
	Id int `json:"id"`
	Node_id string `json:"node_id"`
	Avatar_url string `json:"avatar_url"`
	Gravatar_id string `json:"gravatar_id"`
	Url string `json:"url"`
	Html_url string `json:"html_url"`
	Followers_url string `json:"followers_url"`
	Following_url string `json:"following_url"`
	Gists_url string `json:"gists_url"`
	Starred_url string `json:"starred_url"`
	Subscriptions_url string `json:"subscriptions_url"`
	Organizations_url string `json:"organizations_url"`
	Repos_url string `json:"repos_url"`
	Events_url string `json:"events_url"`
	Received_events_url string `json:"received_events_url"`
	Type string `json:"type"`
	Site_admin bool `json:"site_admin"`
	Name string `json:"name"`
	Company string `json:"company"`
	Blog string `json:"blog"`
	Location string `json:"location"`
	Email string `json:"email"`
	Hireable bool `json:"hireable"`
	Bio string `json:"bio"`
	Public_repos int `json:"public_repos"`
	Public_gists int `json:"public_gists"`
	Followers int `json:"followers"`
	Following int `json:"following"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
}
//Github UserID -》GithubInfo
type GithubUserMap map[string]*GithubInfo


var GithubUser GithubUserMap

// name->id
type noUserError struct {
	userId string
}

// name->id
func (this noUserError) Error() string {
	return "No such user" + this.userId
}

func init() {
	fmt.Println("Controller initialized!")
	GithubUser = make(GithubUserMap)
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

func getUserJson(access_token string) UserJson {
	var url_2 string = "https://api.github.com/user?" + "access_token=" + access_token

	client_2 := &http.Client{}
	response_2, _ := client_2.Get(url_2)
	defer response_2.Body.Close()
	body_2, err_2 := ioutil.ReadAll(response_2.Body)
	if err_2 != nil {
		panic(err_2)
	}

	// 获取ID
	var name string = strings.Split(strings.Split(string(body_2), ",")[0], "\"")[3]
	var uri string = strings.Split(strings.Split(string(body_2), ",")[3], "\"")[3]
	var id string = strings.Split(strings.Split(string(body_2), ",")[1], ":")[1]

	//select id according to name

	data := Data{
		Name: name,
		Uri:  uri,
		Id:   id,
	}
	json := UserJson{
		ErrorCode: 0,
		Data:      data,
	}
	// 结构体加ID字段

	return json
}

func getAccessToken(code string) (accessToken string, err error) {
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
// 参数name->id
func (this GithubUserMap) setGithubUserAccessToken(id string, name string, accessToken string) {
	if _, ok := this[id]; !ok {
		this[id] = new(GithubInfo)
	}
	this[id].AccessToken = accessToken
	this[id].GithubId = id
	this[id].GithubName = name
}

// 参数name-》id
func (this GithubUserMap) getGithubUserAccessToken(userId string) (string, error) {
	if userInfo, ok := this[userId]; ok {
		return userInfo.AccessToken, nil
	} else {
		err := noUserError{userId: userId}
		return "", err
	}
}

//查询项目url是否合法，且判断用户是否有权限导入

func CheckGithubRepoUrl(userId, url string) error {
	_, repoName, err := models.ParseGithubHTTPSUrl(url)
	//TODO:err处理等待解析函数pr合并后更新
	if err != nil {
		return err
	}
	info := GithubUser[userId]
	userName := info.GithubName
	apiUrl := "https://api.github.com/repos/" + userName + "/" + repoName
	resp, err := http.Get(apiUrl)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		if resp.StatusCode == 404 {
			//fmr.Errorf()可直接返回error类型，参数为error.Error()返回值
			return fmt.Errorf("this repo Url is not valid")
		} else {
			return fmt.Errorf("err %d", resp.StatusCode)
		}
	}
	return nil
}

//getWebhooksUrl 可以通过
func registerGithubWebhooks(userId string, repoName string) {
	accessToken, _ := GithubUser.getGithubUserAccessToken(userId)
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
func GetGithubId(username string)(int){
	api:="https://api.github.com/users/"+username
	client := &http.Client{}
	response, _ := client.Get(api)
	defer response.Body.Close()
	body, err_1 := ioutil.ReadAll(response.Body)
	if err_1 != nil {
		panic(err_1)
	}
	var res API_User
	json.Unmarshal(body,&res)
	fmt.Println(res.Id)
	return res.Id
}