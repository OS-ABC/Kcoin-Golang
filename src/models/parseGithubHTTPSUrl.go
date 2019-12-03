package models

import (
	"strings"
)

func ParseGithubHTTPSUrl(url string) (userName string, userRepo string, ErrorCode int) {
	splitedUrl := strings.Split(url, "/")
	userName = splitedUrl[3]
	userRepo = strings.Split(splitedUrl[4], ".")[0] //处理url最后可能出现的.git
	ErrorCode = 0
	return userName, userRepo, ErrorCode
}
