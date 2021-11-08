package router

import (
	"app_upgrade/model"
	"app_upgrade/service"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddRuleHandle(c *gin.Context) {
	var rule model.Rule
	//TODO
	if err := c.ShouldBind(&rule); err != nil {
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
	rule_id_str, exist = c.GetPostForm("ruleid")
	if exist {
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
	rule_id_str, exist = c.GetPostForm("ruleid")
	if exist {
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
	rule_id_str, exist = c.GetPostForm("ruleid")
	if exist {
		var rule_id_int int
		rule_id_int, _ = strconv.Atoi(rule_id_str)
		rule_id = model.RuleIdType(rule_id_int)
		service.RecoverRule(rule_id)
	}
}

// @title             UpgradeCheckHandle
// @description       根据客户端发送的JSON数据，匹配升级包并以JSON格式回传到客户端
// @auth              卢品洲         2021/11/8
// @param             c             请求句柄
func UpgradeCheckHandle(c *gin.Context) {
	var cd model.ClientData
	if err := c.ShouldBind(&cd); err != nil {
		c.String(http.StatusBadRequest, "Input error!")
	}
	hit, pkg_data := service.UpgradeCheck(&cd)
	if hit {
		c.JSON(200, pkg_data)
	} else {
		c.String(http.StatusNoContent, "No upgrade package was matched.")
	}
}

//ghy实现版本
// func UpdateCheckHandle(c *gin.Context) {
// 	var clientData model.ClientData
// 	if err := c.ShouldBind(&clientData); err == nil {
// 		if isHit, pkgData := service.UpdateCheck(clientData); isHit {
// 			c.JSON(http.StatusOK, pkgData)
// 		}
// 	} else {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 	}
// }
