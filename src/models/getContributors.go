package models
import (
	"fmt"
	"io/ioutil"
	"net/http"
)

//获取项目贡献者信息的接口，返回json为贡献者信息
func getContributors(userName string,programName string) string {
	var url_1 string = "https://api.github.com/repos/"+userName+"/"+programName+"/"+"contributors"

	client :=&http.Client{}
	response,_:=client.Get(url_1)
	defer response.Body.Close()
	body,err_1:=ioutil.ReadAll(response.Body)
	if err_1 != nil{
		panic(err_1)}

	fmt.Println(string(body))
	return string(body)
}


