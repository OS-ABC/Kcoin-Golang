package models

import (
	_ "encoding/json"
	"github.com/astaxie/beego"

	//"github.com/astaxie/beego"
	"io/ioutil"
	"net/http"
	"strings"
)

type Data struct{
	Name string `json:"userName"`
	Uri string `json:"headShotUrl"`

}

type Json struct{
	ErrorCode int
	Data Data
}
func GetGithubAuthJson(code string) Json{
	client_id := beego.AppConfig.String("client_id")
	client_secret := beego.AppConfig.String("client_secret")
	var url_1 string="https://github.com/login/oauth/access_token"+"?code="+string(code)+"&client_id="+client_id+"&client_secret="+client_secret

	client :=&http.Client{}
	response,_:=client.Get(url_1)
	defer response.Body.Close()
	body,err_1:=ioutil.ReadAll(response.Body)
	if err_1 != nil{
		panic(err_1)
	}
	var access_token string= strings.Split(strings.Split(string(body),"&")[0],"=")[1]
	var url_2 string="https://api.github.com/user?"+"access_token="+access_token

	client_2:=&http.Client{}
	response_2,_:=client_2.Get(url_2)
	defer response_2.Body.Close()
	body_2,err_2:=ioutil.ReadAll(response_2.Body)
	if err_2!=nil{
		panic(err_2)
	}
	var name string=strings.Split(strings.Split(string(body_2),",")[0],"\"")[3]
	var uri string=strings.Split(strings.Split(string(body_2),",")[3],"\"")[3]

	data :=Data{
		Name:name,
		Uri:uri,
	}
	json :=Json{
		ErrorCode:0,
		Data: data,
	}

	return json
}
