package service

import (
	"app_upgrade/model"
	"fmt"
	"strings"

	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
	_ "gorm.io/gorm"
)

//数据库配置
const (
	userName = "root"
	password = "fanyang0601"
	ip       = "127.0.0.1"
	port     = "3306"
	dbName   = "app_upgrade"
)

func ConnectDb() *gorm.DB {
	path := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8"}, "")

	db, err := gorm.Open(mysql.Open(path), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	//defer db.Close()

	//if err := db.Ping(); err!=nil{
	//	fmt.Println("连接失败")
	//	panic(err)
	//}
	fmt.Println("连接成功")
	return db
}

func SaveRuleToDB(rule *model.Rule) bool {
	//TODO
	var db = ConnectDb()
	db.Create(&rule)

	return true
}

func GetRuleFromDB(rule_id model.RuleIdType) model.Rule {
	//TODO
	var db = ConnectDb()
	var rule model.Rule
	db.First(&rule, "rid = ?", rule_id)
	return rule
}

func GetRulesFromDB() (rules []model.Rule) {
	//TODO
	var db = ConnectDb()
	db.Find(&rules)
	return rules
}

func PullDeviceIdListFromDB(rule_id model.RuleIdType) (device_id_list string) {
	//TODO
	var db = ConnectDb()
	var rule model.Rule
	db.First(&rule, "rid = ?", rule_id)
	return rule.DeviceIdList
}

// @title             DeleteRuleFromDB
// @description       从数据库删除规则
// @auth              刘晶玉         2021/11/14
// @param             rule_id          规则id
func DeleteRuleFromDB(rule_id model.RuleIdType) bool {
	var db = ConnectDb()
	//查询是否存在
	var rule model.Rule
	if db.First(&rule, "rid = ?", rule_id).Error != nil{
		return false
	}
	//删除
	db.Delete(&model.Rule{},"rid=?",rule_id)
	return true
}

// @title             PauseRuleFromDB
// @description       从数据库暂停规则
// @auth              刘晶玉         2021/11/14
// @param             rule_id          规则id
func PauseRuleFromDB(rule_id model.RuleIdType) bool {
	var db = ConnectDb()
	var rule model.Rule
	//查询存在
	if db.First(&rule, "rid = ?", rule_id).Error != nil{
		return false
	}
	//更新
	db.Model(model.Rule{}).Where("state", true).Update("state", false)
	return true
}

// @title             RecoverRuleFromDB
// @description       从数据库恢复规则
// @auth              刘晶玉         2021/11/14
// @param             rule_id          规则id
func RecoverRuleFromDB(rule_id model.RuleIdType) bool {
	var db = ConnectDb()
	var rule model.Rule
	//查询存在
	if db.First(&rule, "rid = ?", rule_id).Error != nil{
		return false
	}
	//更新
	db.Model(model.Rule{}).Where("state", false).Update("state", true)
	return true
}
