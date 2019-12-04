package controllers

import (
	"bytes"
	"context"
	_ "encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
	"net/http"
	"strings"
	"github.com/astaxie/beego/orm"

)

/**
 * 这是一个全局数据结构，目前只有两个字段，用来保存Github名和对应access_token
 */
type GithubInfo struct {
	GithubName  string
	AccessToken string
}

type GithubUserMap map[string]*GithubInfo

var GithubUser GithubUserMap

type noUserError struct {
	userName string
}

func (this noUserError) Error() string {
	return "No such user" + this.userName
}

func init() {
	fmt.Println("Controller initialized!")
	GithubUser = make(GithubUserMap)
}

type Data struct {
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
	var name string = strings.Split(strings.Split(string(body_2), ",")[0], "\"")[3]
	var uri string = strings.Split(strings.Split(string(body_2), ",")[3], "\"")[3]

	data := Data{
		Name: name,
		Uri:  uri,
	}
	json := UserJson{
		ErrorCode: 0,
		Data:      data,
	}

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
func (this GithubUserMap) setGithubUserAccessToken(name string, accessToken string) {
	if _, ok := this[name]; !ok {
		this[name] = new(GithubInfo)
	}
	this[name].AccessToken = accessToken
	this[name].GithubName = name
}

func (this GithubUserMap) getGithubUserAccessToken(name string) (string, error) {
	if userInfo, ok := this[name]; ok {
		return userInfo.AccessToken, nil
	} else {
		err := noUserError{userName: name}
		return "", err
	}
}

func CheckGithubRepoUrl(userId , url string) error{
	userName, repoName, err := ParseGithubUrl(url)
	//TODO:err处理等待解析函数pr合并后更新
	if err != nil {
		return err
	}
	apiUrl := "https://api.github.com/repos/"+userName+"/"+repoName
	resp, err := http.Get(apiUrl)
	if err != nil {
		return err
	}
	if resp.StatusCode == 404 {
		//fmr.Errorf()可直接返回error类型，参数为error.Error()返回值
		return fmt.Errorf("this repo Url is not valid")
	}
	sqlQueryUserName := `SELECT user_name FROM "K_User" WHERE github_user_name = ?`
	o := orm.NewOrm()
	o.Using("defalut")
	var thisUserName string
	o.Raw(sqlQueryUserName, userId).QueryRow(&thisUserName)
	if thisUserName == userName {
		return nil
	} else {
		return fmt.Errorf("the user does not own this repo")
	}

}

func registerGithubWebhooks(userName string, repoName string) {
	accessToken, _ := GithubUser.getGithubUserAccessToken(userName)
	postPayload := getPayloadOfRegisterGithubWebhooks()
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
