package models

import (
	"strings"
)

func ParseGithubHTTPSUrl(url string) (userName string, userRepo string, ErrorCode error) {
	splitedUrl := strings.Split(url, "/")
	userName = splitedUrl[3]
	userRepo = strings.Split(splitedUrl[4], ".")[0] //处理url最后可能出现的.git
	ErrorCode = nil                                 //还未定义错误类型，先返回nil
	return userName, userRepo, ErrorCode
}
