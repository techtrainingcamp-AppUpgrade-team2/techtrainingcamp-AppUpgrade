package service

import (
	"app_upgrade/model"
	"log"
)

func AddRule(rule *model.Rule) bool {
	//TODO
	//rid唯一标识
	currWoker := &idworker.IdWorker{}
	currWoker.InitIdWorker(1000, 1)
	newID, err := currWoker.NextId()
	if err != nil {
		log.Fatal(err.Error())
		return false
	}
	rule.Rid = model.RuleIdType(newID)
	//state
	rule.State = true
	//存储到数据库
	if SaveRuleToDB(rule) == false {
		return false
	}
	//添加到cache白名单
	AddDeviceIDList(rule.Rid, rule.DeviceIdList)
	return true
}

func DeleteRule(rule_id model.RuleIdType) bool {
	//TODO
	DeleteRuleFromDB(rule_id)
	//应该还有删除cache对应DeviceId白名单，且其他rid没有对应DeviceId，以下两个函数同理

	return true
}

func PauseRule(rule_id model.RuleIdType) bool {
	//TODO
	return PauseRuleFromDB(rule_id)
}

func RecoverRule(rule_id model.RuleIdType) bool {
	//TODO
	return RecoverRuleFromDB(rule_id)
}
