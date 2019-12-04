package models

import "testing"

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
		{args:args{userName:"rjkris", programName:"fluffy-robot"}, want:"rjkris "},//PASS
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getContributors(tt.args.userName, tt.args.programName); got != tt.want {
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
		{args:args{url:"https://github.com/OS-ABC/HelloWorld"}, want:115},//FAIL 通过该API获得的contributors可能不全
		{args:args{url:"https://github.com/OS-ABC/Kcoin-Golang"}, want:17},//PASS
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetContributorNum(tt.args.url); got != tt.want {
				t.Errorf("GetContributorNum() = %v, want %v", got, tt.want)
			}
		})
	}
}