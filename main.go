package main

import (
	_ "Kcoin-Golang/conf"
	"Kcoin-Golang/routers"
)

func main() {
	// 初始化路由
	r := routers.RouterInit()
	r.Run(":8080")
}
