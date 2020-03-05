// TODO 该文件并入service层的github_helper.go中, 并删除main函数, 改为测试覆盖.
package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type GithubRepos []struct {
	Name string `json:"name"`
	Url  string `json:"html_url"`
}

//type ReposInfo struct {
//	ErrorCode string `json:"error_code"`
//	Data GithubRepos `json:"data"`
//}
//获取一个用户在github上所有公开的repository，返回json
func GetGithubRepos(user string) (string, error) {
	//var reposInfo ReposInfo
	//reposInfo.ErrorCode = "default Error"
	var url string = "https://api.github.com/users/" + user + "/repos"

	client := &http.Client{}
	response, _ := client.Get(url)
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	var repos GithubRepos
	var projects ProjectInfo
	err1 := json.Unmarshal([]byte(body), &repos)
	if err1 != nil {
		panic(err1)
	}
	for i := range repos {
		p := &Project{}
		p.ProjectUrl = repos[i].Url
		p.ProjectName = repos[i].Name
		projects.Data = append(projects.Data, p)
	}
	//reposInfo.ErrorCode = "0"
	//reposInfo.Data = repos
	//res,err2 := json.Marshal(&reposInfo)
	res, err2 := json.Marshal(&projects)
	if err2 != nil {
		panic(err2)
	}
	return string(res), nil
}

//测试输出
func main() {
	fmt.Println(GetGithubRepos("Darkone0"))
}
