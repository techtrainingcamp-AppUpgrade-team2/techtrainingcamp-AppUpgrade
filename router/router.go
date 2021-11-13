package router

import (
	"app_upgrade/service"

	"github.com/gin-gonic/gin"
)

//GET请求路由表
var router_get = map[string]func(*gin.Context){
	"/ping":    service.Pong,
	"/upgrade": UpgradeCheckHandle,
	"/get":     GetRulesHandle,
}

//POST请求路由表
var router_post = map[string]func(*gin.Context){
	"rule/add":     AddRuleHandle,
	"rule/delete":  DeleteRuleHandle,
	"rule/pause":   PauseRuleHandle,
	"rule/recover": RecoverRuleHandle,
}

//...其它路由表

// @title             RouterInit
// @description       初始化router，并返回
// @auth              卢品洲         2021/11/7
// @param             -
// @return            -             router
func RouterInit() *gin.Engine {
	router := gin.Default()

	//配置路由
	for path, handle := range router_get {
		router.GET(path, handle)
	}
	for path, handle := range router_post {
		router.POST(path, handle)
	}
	return router
}
