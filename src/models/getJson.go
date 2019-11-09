package models

import (
	_ "encoding/json"
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
func GetJson(code string) Json{
	client_id := "c698bd09d35414343da4"
	client_secret :="e680c1f3ae01fab2175efa4b3366cc43d7ca36ac"
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
