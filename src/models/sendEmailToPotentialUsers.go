// TODO 该文件不应该放在model层, 应该放在service层
package models

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/astaxie/beego/utils"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"time"
)

//用于登录时验证
type AuthenticityToken struct {
	Token string
}

//同样用于登录时验证
type TimeStamp struct {
	Time   string
	Secret string
}

//http请求
type App struct {
	Client *http.Client
}

const (
	baseURL       = "https://github.com"
	githubUserApi = "https://api.github.com/users/"
)

var ( //用于模拟登录，个人账号，请勿登录
	username = "kcoinTest"
	password = "Kcoin123456."
)

//用于接收爬取到的邮箱
type Email struct {
	Email string
}

type UsersApiJson struct {
	Login             string      `json:"login"`
	ID                int         `json:"id"`
	NodeID            string      `json:"node_id"`
	AvatarURL         string      `json:"avatar_url"`
	GravatarID        string      `json:"gravatar_id"`
	URL               string      `json:"url"`
	HTMLURL           string      `json:"html_url"`
	FollowersURL      string      `json:"followers_url"`
	FollowingURL      string      `json:"following_url"`
	GistsURL          string      `json:"gists_url"`
	StarredURL        string      `json:"starred_url"`
	SubscriptionsURL  string      `json:"subscriptions_url"`
	OrganizationsURL  string      `json:"organizations_url"`
	ReposURL          string      `json:"repos_url"`
	EventsURL         string      `json:"events_url"`
	ReceivedEventsURL string      `json:"received_events_url"`
	Type              string      `json:"type"`
	SiteAdmin         bool        `json:"site_admin"`
	Name              string      `json:"name"`
	Company           interface{} `json:"company"`
	Blog              string      `json:"blog"`
	Location          string      `json:"location"`
	Email             string      `json:"email"`
	Hireable          interface{} `json:"hireable"`
	Bio               interface{} `json:"bio"`
	PublicRepos       int         `json:"public_repos"`
	PublicGists       int         `json:"public_gists"`
	Followers         int         `json:"followers"`
	Following         int         `json:"following"`
	CreatedAt         time.Time   `json:"created_at"`
	UpdatedAt         time.Time   `json:"updated_at"`
}

/*
func (app *App) GetToken() AuthenticityToken
该函数用于获取github的登录时的authenticity_token，是一个隐藏的值，相当于验证码
*/
func (app *App) GetToken() (AuthenticityToken, TimeStamp) {
	var authenticityToken AuthenticityToken
	var timeAndsecret TimeStamp
	loginURL := baseURL + "/login"
	client := app.Client
	response, err := client.Get(loginURL)
	if err != nil {
		log.Fatalln("Error fetching response. ", err)
		return authenticityToken, timeAndsecret
	}
	defer response.Body.Close()

	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal("Error loading HTTP response body. ", err)
	}

	token, exist := document.Find("input[name='authenticity_token']").Attr("value")
	if exist != true {
		log.Fatal("Error finding authenticity_token,does not exist", exist)
		exist = true
	}

	fmt.Println("token is", token)
	authenticityToken.Token = token
	time, exist := document.Find("input[name='timestamp']").Attr("value")
	if exist != true {
		log.Fatal("Error finding timestamp,does not exist", exist)
		exist = true
	}
	secret, exist := document.Find("input[name='timestamp_secret']").Attr("value")
	if exist != true {
		log.Fatal("Error finding timestamp,does not exist", exist)
		exist = true
	}
	fmt.Println("timestamp is ", time)
	fmt.Println("timestamp_secret is ", secret)
	timeAndsecret.Secret = secret
	timeAndsecret.Time = time
	return authenticityToken, timeAndsecret
}

/*
func (app *App) Login()
该函数用于模拟登录，如果不登录，爬取到的内容是“sign up to view email”
*/
func (app *App) Login() {
	client := app.Client
	//通过GetToken（）获得登录所需的动态信息
	authenticityToken, timeAndsecret := app.GetToken()
	loginURL := baseURL + "/session"
	//模拟登录的文件header中的Form data,详情请看github登录时候的发包
	data := url.Values{
		"utf8":                    {"✓"},
		"commit":                  {"Sign in"},
		"ga_id":                   {"76806099.1573434977"},
		"login":                   {username},
		"password":                {password},
		"webauthn-support":        {"supported"},
		"webauthn-iuvpaa-support": {"unsupported"},

		"required_field_440e": {},
		"timestamp":           {timeAndsecret.Time},
		"timestamp_secret":    {timeAndsecret.Secret},
		"authenticity_token":  {authenticityToken.Token},
	}
	//模拟post请求，发送header
	response, err := client.PostForm(loginURL, data)
	if err != nil {
		fmt.Println(err)
		log.Fatalln(err)
	}
	defer response.Body.Close()
	_, err = ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("status code is ", response.StatusCode)
	fmt.Println("login done")
}

/*
func (app *App) getEmail(user,githubToken string) Email
功能：获得用户user的邮箱
参数：user github用户名
	 githubToken 此时登录系统的github账号的githubToken
返回：user对应的邮箱（用户未公布邮箱，未测试）
*/
func (app *App) getEmail(user, githubToken string) Email {
	//个人主页，如果邮箱公开，那么会显示在主页
	personalURL := githubUserApi + user + "?access_token=" + githubToken

	client := &http.Client{}
	response, _ := client.Get(personalURL)
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	var userApiJson UsersApiJson
	err = json.Unmarshal(body, &userApiJson)
	if err != nil {
		log.Fatal("Error Unmarshaling json response body from githubUserApi. ", err)
		panic(err)
	}
	return Email{Email: userApiJson.Email}
}

/*
func SendEMailToPotentialUsers(users []string)([]string ,error)
功能：将邮件发送给users中的每一个人
参数：users 用户列表，github用户名的切片，待发送邮件的名单
	githubToken 此时登录系统的github账号的githubToken
返回：邮箱不对外显示，因此无法发送邮件的成员用户名集合
*/
func SendEMailToPotentialUsers(users []string, githubToken string) ([]string, error) {
	jar, _ := cookiejar.New(nil)
	app := App{
		Client: &http.Client{Jar: jar},
	}
	app.Login()

	//未能成功发送邮件的用户名单
	var UsersNotSend []string
	for i := 0; i < len(users); i++ {
		email := app.getEmail(users[i], githubToken)
		if email.Email == "" {
			//如果得到邮箱，给他加到列表里，最后返回
			UsersNotSend = append(UsersNotSend, users[i])
		} else {
			//如果得到了邮箱
			config :=
				`{"username":"kcoin_golang@163.com","password":"kcoin163","host":"smtp.163.com","port":25}`
			// 通过存放配置信息的字符串，创建Email对象
			temail := utils.NewEMail(config)
			// 指定邮件的基本信息
			temail.To = []string{email.Email}           //指定收件人邮箱地址
			temail.From = "kcoin_golang@163.com"        //指定发件人的邮箱地址，这是我注册的kcoin项目邮箱，账号kcoin_golang@163.com 密码kcoingolang
			temail.Subject = "Kcoin：Welcome to join us" //指定邮件的标题
			//这是邮件的内容
			temail.HTML = `<html>
                    <head>
                    </head>
                        <body>
                         <div>你好，我们是开原激励社区Kcoin，现在你所在的github项目已经加入Kcoin，我们同样欢迎您的加入</div>
                       </body>
                    </html>`
			// 发送邮件
			err := temail.Send()
			if err != nil {
				return UsersNotSend, err
			}
		}
	}
	return UsersNotSend, nil
}
