package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"kcoin-golang/src/models"
	_ "kcoin-golang/src/routers"
)

func main() {
	var names = make([]string, 2)
	names[0] = "test"
	names[1] = "test2"

	result := models.GetUserInfo(names[1])
	fmt.Println(result)
}

