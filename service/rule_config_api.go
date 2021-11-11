package service

import (
	"app_upgrade/model"
	"log"

	idworker "github.com/gitstliu/go-id-worker"
)

// @title             AddRule
// @description       增加规则
// @auth              刘晶玉         2021/11/11
// @param             rule             规则
func AddRule(rule *model.Rule) bool {
	//TODO
	//生成rid唯一标识
	currWoker := &idworker.IdWorker{}
	currWoker.InitIdWorker(1000, 1)
	newID, err := currWoker.NextId()
	if err != nil {
		log.Fatal(err.Error())
		return false
	}
	rule.Rid = model.RuleIdType(newID)
	//state为可用
	rule.State = true
	//存储到数据库
	if SaveRuleToDB(rule) == false {
		return false
	}
	//添加到cache白名单
	AddDeviceIDList(rule.Rid, rule.DeviceIdList)
	return true
}

// @title             DeleteRule
// @description       删除规则
// @auth              刘晶玉         2021/11/11
// @param             rule_id             规则id
func DeleteRule(rule_id model.RuleIdType) bool {
	//TODO
	DeleteRuleFromDB(rule_id)	//数据库删除rule
	RemoveDeviceIDList(rule_id)	//删除cache对应DeviceId白名单
	return true
}

// @title             PauseRule
// @description       暂停规则
// @auth              刘晶玉         2021/11/11
// @param             rule_id             规则id
func PauseRule(rule_id model.RuleIdType) bool {
	//TODO
	PauseRuleFromDB(rule_id)	//数据库修改rule的状态state
	RemoveDeviceIDList(rule_id)	//删除cache对应DeviceId白名单
	return true
}

// @title             RecoverRule
// @description       恢复规则
// @auth              刘晶玉         2021/11/11
// @param             rule_id             规则id
func RecoverRule(rule_id model.RuleIdType) bool {
	//TODO
	RecoverRuleFromDB(rule_id)	//数据库修改rule的状态state
	rule := GetRuleFromDB(rule_id)	//从数据库取得rid对应rule详细
	AddDeviceIDList(rule.Rid, rule.DeviceIdList)	//添加到cache白名单
	return true
}

// @title             GetRules
// @description       得到当前正在生效和暂停的规则
// @auth              刘晶玉         2021/11/11
func GetRules() (rules []model.Rule) {
	//TODO
	rules = GetRulesFromDB() //获取所有规则
	return rules
}