package models

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type GithubRepos []struct {
	Name string `json:"name"`
	Url string `json:"html_url"`
}

type ReposInfo struct {
	ErrorCode string `json:"error_code"`
	Data GithubRepos `json:"data"`
}
//获取一个用户在github上所有公开的repository，返回json
func getGithubRepos(user string)(string, error){
	var reposInfo ReposInfo
	reposInfo.ErrorCode = "default Error"
	var url string = "https://api.github.com/users/" + user + "/repos"

	client := &http.Client{}
	response,_ := client.Get(url)
	defer response.Body.Close()
	body,err := ioutil.ReadAll(response.Body)
	if err != nil{
		panic(err)
	}
	var repos GithubRepos
	err1 := json.Unmarshal([]byte(body), &repos)
	if err1 != nil{
		panic(err1)
	}
	reposInfo.ErrorCode = "0"
	reposInfo.Data = repos
	res,err2 := json.Marshal(&reposInfo)
	if err2 != nil{
		panic(err2)
	}
	return string(res), nil
}

//测试输出
//func main()  {
//	fmt.Println(getGithubRepos("OS-ABC"))
//}
