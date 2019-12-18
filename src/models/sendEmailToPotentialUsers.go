// TODO 该文件不应该放在model层, 应该放在service层
package models

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/astaxie/beego/utils"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
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
	baseURL = "https://github.com"
)

var ( //用于模拟登录，个人账号，请勿登录
	username = "kcoinTest"
	password = "Kcoin123456."
)

//用于接收爬取到的邮箱
type Email struct {
	Email string
}

var res_error error
var UsersNotSend []string

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
func (app *App) getEmails(users string) []Email
使用goquery来对得到的html进行解析，最后解析出页面中的邮箱信息，返回
*/
func (app *App) getEmails(users string) []Email {
	//个人主页，如果邮箱公开，那么会显示在主页
	personalURL := baseURL + "/" + users
	client := app.Client

	response, err := client.Get(personalURL)
	if err != nil {
		log.Fatalln("Error fetching response. ", err)
	}
	defer response.Body.Close()
	//建立一个新的goquery对象
	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal("Error loading HTTP response body. ", err)
	}
	//emails为返回信息
	var emails []Email
	//在class为js-profile-editable-area的，下面的ul，下面的li中寻找a标签，因为邮箱放在一个a标签中
	document.Find(".js-profile-editable-area ul li").Each(func(i int, s *goquery.Selection) {
		//获取该标签中的内容
		email := s.Find("a").Text()
		emailEntry := Email{
			Email: email,
		}
		emails = append(emails, emailEntry)
	})

	return emails
}

/*
func SendEMailToPotentialUsers(users []string)([]string ,error)
该函数功能是通过username，在网页上爬取邮箱，然后发送内容给他们
参数users为username的切片
*/
func SendEMailToPotentialUsers(users []string) ([]string, error) {
	jar, _ := cookiejar.New(nil)
	app := App{
		Client: &http.Client{Jar: jar},
	}
	app.Login()
	//返回信息

	for i := 0; i < len(users); i++ {
		emails := app.getEmails(users[i])
		length := len(emails)
		fmt.Printf("the user %s emails has %d entry \n", users[i], length)
		if length == 0 {
			fmt.Println("this user's page have no public email")
			//如果没爬到邮箱，给他加到列表里，最后返回
			UsersNotSend = append(UsersNotSend, users[i])
		} else {
			//如果爬到了邮箱
			//这里为什么用emails[0]呢，因为某种玄学原因，emails里有两个重复的邮箱，我们用一个就行了
			fmt.Println("now sending email")
			currentemail := emails[0].Email
			fmt.Println("email is :",currentemail)
			//for index, email := range emails {
			//    fmt.Printf("%d: %s\n", index+1, email.Email)
			//}
			//`{"username":"邮箱名称","password":"163的token","host":"SMTP服务器地址","port":对应端口号}`
			config :=
				`{"username":"kcoin_golang@163.com","password":"kcoin163","host":"smtp.163.com","port":25}`
			// 通过存放配置信息的字符串，创建Email对象
			temail := utils.NewEMail(config)
			// 指定邮件的基本信息
			temail.To = []string{currentemail}          //指定收件人邮箱地址
			temail.From = "kcoin_golang@163.com"        //指定发件人的邮箱地址，这是我注册的kcoin项目邮箱，账号kcoin_golang@163.com 密码kcoingolang
			temail.Subject = "Kcoin：Welcome to join us" //指定邮件的标题
			//这是邮件的内容
			temail.HTML = `<html>
                    <head>
                    </head>
                        <body>
                         <div>你好，我们是开原激励社区Kcoin，现在你所在的github项目已经加入Kcoin，我们同样欢迎您的加入</div>
                       </body>
                    </html>` //指定邮件内容
			// 发送邮件
			err := temail.Send()
			res_error = err
			if err != nil {
				return UsersNotSend, res_error
			}
		}
	}
	fmt.Println(UsersNotSend)
	return UsersNotSend, res_error
}
