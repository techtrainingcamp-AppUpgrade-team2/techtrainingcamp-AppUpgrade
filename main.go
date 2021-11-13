package main

import (
	"app_upgrade/router"
	"app_upgrade/service"
)

func main() {
	router := router.RouterInit()
	router.Run(":9090")
	//testDb()
}

func testDb() {

	//test for data_base.go
	//rule.Rid = 1
	//service.SaveRuleToDB(&rule)

	//test for cache.go
	service.AddDeviceIDList(1, service.PullDeviceIdListFromDB(1))
	println(service.CheckDeviceIDListList(1, "hello"))
	println(service.CheckDeviceIDListList(1, "he2o"))
}
