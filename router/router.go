package router

import (
	"app_upgrade/service"

	"github.com/gin-gonic/gin"
)

//GET请求路由表
var router_get = map[string]func(*gin.Context){
	"/":       service.Demo,
	"upgrade": UpgradeCheckHandle,
}

//POST请求路由表
var router_post = map[string]func(*gin.Context){
	"rule/add": AddRuleHandle,

	"rule/delete":  DeleteRuleHandle,
	"rule/pause":   PauseRuleHandle,
	"rule/recover": RecoverRuleHandle,
}

//...其它路由表

func RouterInit() *gin.Engine {
	router := gin.Default()

	//配置路由
	for path, handle := range router_get {
		router.GET(path, handle)
	}
	for path, handle := range router_post {
		router.POST(path, handle)
	}
	//...

	return router
}
