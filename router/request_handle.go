package router

import (
	"app_upgrade/model"
	"app_upgrade/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// func AddRuleHandle(c *gin.Context) {
// 	var rule model.Rule
// 	//TODO
// 	service.AddRule(&rule)
// }

// func DeleteRuleHandle(c *gin.Context) {
// 	var rule_id model.RuleIdType
// 	//TODO
// 	service.DeleteRule(rule_id)
// }

// func PauseRuleHandle(c *gin.Context) {
// 	var rule_id model.RuleIdType
// 	//TODO
// 	service.DeleteRule(rule_id)
// }

// func RecoverRuleHandle(c *gin.Context) {
// 	var rule_id model.RuleIdType
// 	//TODO
// 	service.DeleteRule(rule_id)
// }

func UpdateCheckHandle(c *gin.Context) {
	var clientData model.ClientData
	if err := c.ShouldBind(&clientData); err == nil {
		if isHit, pkgData := service.UpdateCheck(clientData); isHit {
			c.JSON(http.StatusOK, pkgData)
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
