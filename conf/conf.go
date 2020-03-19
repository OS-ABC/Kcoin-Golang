package conf

import (
	"gopkg.in/ini.v1"
)

// 所有的配置信息
type AllConfig struct {
	MySQL
	GithubOAuth
}

// MySQL配置信息
type MySQL struct {
	Host     string `ini:"host"`
	Port     string `ini:"port"`
	User     string `ini:"user"`
	DBname   string `ini:"dbname"`
	Password string `ini:"password"`
}

// GithubOAuth配置信息
type GithubOAuth struct {
	ClientID     string `ini:"clientID"`
	ClientSecret string `ini:"clientSecret"`
}

// 保存读取的配置信息的变量
var Config AllConfig

func init() {
	err := LoadConfig()
	if err != nil {
		panic(err)
	}
}

func LoadConfig() error {
	err := ini.MapTo(&Config, "./conf/config.ini")
	if err != nil {
		return err
	}
	return nil
}
