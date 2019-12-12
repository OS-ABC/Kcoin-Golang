package models

import (
	"fmt"
	"strings"
)

func ParseGithubHTTPSUrl(url string) (userName string, userRepo string,  err error) {
	splitedUrl := strings.Split(url, "/")
	if len(splitedUrl) != 5 {return "", "", fmt.Errorf("invalid url")}
	userName = splitedUrl[3]
	userRepo = strings.Split(splitedUrl[4], ".")[0] //处理url最后可能出现的.git
	return userName, userRepo, nil
}
