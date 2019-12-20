package service

import (
	"fmt"
	"net/http"
	"strings"
)

func ParseGithubHTTPSUrl(url string) (userName string, userRepo string, err error) {
	splitedUrl := strings.Split(url, "/")
	if len(splitedUrl) != 5 {
		return "", "", fmt.Errorf("invalid url")
	}
	userName = splitedUrl[3]
	userRepo = strings.Split(splitedUrl[4], ".")[0] //处理url最后可能出现的.git
	return userName, userRepo, nil
}

//查询项目url是否合法，且判断用户是否有权限导入
func CheckGithubRepoUrl(githubName, url string) error {
	_, repoName, err := ParseGithubHTTPSUrl(url)
	//TODO:err处理等待解析函数pr合并后更新
	if err != nil {
		return err
	}
	//info := GithubUser[userId]
	//userName := info.GithubName
	apiUrl := "https://api.github.com/repos/" + githubName + "/" + repoName
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
