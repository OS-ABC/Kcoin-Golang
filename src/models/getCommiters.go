package models 
//api:https://api.github.com/repos/OS-ABC/Kcoin-Golang/commits


import (
	"fmt"
	"io/ioutil"
	"net/http"
)

//验证是否是项目commiter的接口，参数为项目名称和用户名称
func checkAutho(userName string,programName string) string {
//请求：获取此项目的commiter信息
	var url_1 string = "https://api.github.com/repos/"+userName+"/"+programName+"/"+"commits"
	client :=&http.Client{}
	response,_:=client.Get(url_1)
	defer response.Body.Close()
	body,err_1:=ioutil.ReadAll(response.Body)
	if err_1 != nil{
		panic(err_1)}

	fmt.Println(string(body))
	return string(body)
}

//使用实例：
//func main() string{
//	var info string=checkAutho("PHP-is-Best","Kcoin-Golang")
//	fmt.Println(info)
//	return info
//}

