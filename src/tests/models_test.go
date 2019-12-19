package test

import (
	"Kcoin-Golang/src/models"
	"Kcoin-Golang/src/service"
	"testing"
)

func Test_getContributors(t *testing.T) {
	type args struct {
		userName    string
		programName string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		//测试用例
		{args: args{userName: "rjkris", programName: "fluffy-robot"}, want: "rjkris "}, //PASS
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := service.GetContributors(tt.args.userName, tt.args.programName); got != tt.want {
				t.Errorf("getContributors() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetContributorNum(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		//测试用例
		{args: args{url: "https://github.com/OS-ABC/HelloWorld"}, want: 115},  //FAIL 通过该API获得的contributors可能不全
		{args: args{url: "https://github.com/OS-ABC/Kcoin-Golang"}, want: 17}, //PASS
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := service.GetContributorNum(tt.args.url); got != tt.want {
				t.Errorf("GetContributorNum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetStarNum(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		//测试用例
		{args: args{url: "https://github.com/OS-ABC/Kcoin-Golang"}, want: 16}, //PASS
		{args: args{url: "https://github.com/OS-ABC/HelloWorld"}, want: 86},   //PASS
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := service.GetStarNum(tt.args.url); got != tt.want {
				t.Errorf("GetStarNum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseGithubHTTPSUrl(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name         string
		args         args
		wantUserName string
		wantUserRepo string
	}{
		//测试用例
		{args: args{url: "https://github.com/OS-ABC/Kcoin-Golang"}, wantUserName: "OS-ABC", wantUserRepo: "Kcoin-Golang"},   //PASS
		{args: args{url: "https://github.com/zhang2j/HelloWorld.git"}, wantUserName: "zhang2j", wantUserRepo: "HelloWorld"}, //PASS
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotName, gotRepo, _ := service.ParseGithubHTTPSUrl(tt.args.url); gotName != tt.wantUserName || gotRepo != tt.wantUserRepo {
				t.Errorf("userName = %v, want %v; userRepo = %v; want %v", gotName, tt.wantUserName, gotRepo, tt.wantUserRepo)
			}
		})
	}
}

func TestGetProjectsCC(t *testing.T) {
	type args struct {
		projectName string
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		//测试用例
		{args: args{projectName: "baidu"}, want: 1},        //PASS
		{args: args{projectName: "i_love_Kcoin"}, want: 3}, //PASS
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := models.GetProjectsCC(tt.args.projectName); got != tt.want {
				t.Errorf("CCNum = %v, want %v", got, tt.want)
			}
		})
	}
}