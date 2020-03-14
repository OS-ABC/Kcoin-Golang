package main

import (
	"Kcoin-Golang/routers"
)

func main() {
	// 初始化路由
	r := routers.RouterInit()
	r.Run(":8080")
}
