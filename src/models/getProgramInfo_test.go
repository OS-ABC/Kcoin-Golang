package models

import "testing"

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
		{args: args{url:"https://github.com/OS-ABC/Kcoin-Golang"}, want:15},//PASS
		{args: args{url:"https://github.com/OS-ABC/HelloWorld"}, want:86},//PASS
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetStarNum(tt.args.url); got != tt.want {
				t.Errorf("GetStarNum() = %v, want %v", got, tt.want)
			}
		})
	}
}