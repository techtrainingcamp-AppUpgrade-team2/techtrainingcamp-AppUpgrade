package service

import (
	"app_upgrade/model"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
	_ "gorm.io/gorm"
	"fmt"
	"strings"
)

//数据库配置
const (
	userName = "root"
	password = "fanyang0601"
	ip = "127.0.0.1"
	port = "3306"
	dbName = "app_upgrade"
)



func ConnectDb() *gorm.DB {
	path := strings.Join([]string{userName, ":", password, "@tcp(",ip, ":", port, ")/", dbName, "?charset=utf8"}, "")

	db, err := gorm.Open(mysql.Open(path), &gorm.Config{})
	if err != nil{
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
func GetRuleFromDB(rule_id model.RuleIdType) model.Rule{
	//TODO
	var db = ConnectDb()
	var rule model.Rule
	db.First(&rule, "rid = ?",rule_id)
	return rule
}

func PullDeviceIdListFromDB(rule_id model.RuleIdType) (device_id_list string) {
	//TODO
	var db = ConnectDb()
	var rule model.Rule
	db.First(&rule, "rid = ?",rule_id)
	return rule.DeviceIdList
}
