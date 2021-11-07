package router

import (
	"app_upgrade/model"
	"app_upgrade/service"

	"github.com/gin-gonic/gin"
)

func AddRuleHandle(c *gin.Context) {
	var rule model.Rule
	//TODO
	service.AddRule(&rule)
}

func DeleteRuleHandle(c *gin.Context) {
	var rule_id model.RuleIdType
	//TODO
	service.DeleteRule(rule_id)
}

func PauseRuleHandle(c *gin.Context) {
	var rule_id model.RuleIdType
	//TODO
	service.DeleteRule(rule_id)
}

func RecoverRuleHandle(c *gin.Context) {
	var rule_id model.RuleIdType
	//TODO
	service.DeleteRule(rule_id)
}

func UpdateCheckHandle(c *gin.Context) {
	//var cd model.ClientData
	//TODO
	//hit, pkg_data := service.UpgradeCheck(&cd)
	//TODO
}
