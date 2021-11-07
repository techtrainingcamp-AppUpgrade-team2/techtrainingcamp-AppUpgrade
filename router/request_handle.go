package router

import (
	"app_upgrade/model"
	"app_upgrade/service"
	"log"
	"strconv"
	"github.com/gin-gonic/gin"
)

func AddRuleHandle(c *gin.Context) {
	var rule model.Rule
	//TODO
	if err:=c.ShouldBind(&rule);err != nil{
		log.Fatal(err.Error())
		return
	}
	service.AddRule(&rule)
}

func DeleteRuleHandle(c *gin.Context) {
	var rule_id model.RuleIdType
	//TODO
	var rule_id_str string
	var exist bool
	rule_id_str,exist = c.GetPostForm("ruleid")
	if exist{
		var rule_id_int int
		rule_id_int, _ = strconv.Atoi(rule_id_str)
		rule_id = model.RuleIdType(rule_id_int)
		service.DeleteRule(rule_id)
	}
}

func PauseRuleHandle(c *gin.Context) {
	var rule_id model.RuleIdType
	//TODO
	var rule_id_str string
	var exist bool
	rule_id_str,exist = c.GetPostForm("ruleid")
	if exist{
		var rule_id_int int
		rule_id_int, _ = strconv.Atoi(rule_id_str)
		rule_id = model.RuleIdType(rule_id_int)
		service.PauseRule(rule_id)
	}
}

func RecoverRuleHandle(c *gin.Context) {
	var rule_id model.RuleIdType
	//TODO
	var rule_id_str string
	var exist bool
	rule_id_str,exist = c.GetPostForm("ruleid")
	if exist{
		var rule_id_int int
		rule_id_int, _ = strconv.Atoi(rule_id_str)
		rule_id = model.RuleIdType(rule_id_int)
		service.RecoverRule(rule_id)
	}
}

func UpdateCheckHandle(c *gin.Context) {
	//var cd model.ClientData
	//TODO
	//hit, pkg_data := service.UpgradeCheck(&cd)
	//TODO
}
