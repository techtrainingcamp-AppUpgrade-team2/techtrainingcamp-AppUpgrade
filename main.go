package main

import (
	"app_upgrade/router"
)

func main() {
	router := router.RouterInit()
	router.Run()
}
