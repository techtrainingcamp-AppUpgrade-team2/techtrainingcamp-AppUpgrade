package service

import "app_upgrade/model"

//TODO:选型
var DeviceIDCache map[model.RuleIdType]map[model.DeviceIdType]struct{}

//往Cache中添加指定规则的白名单
func AddDeviceIDList(rule_id model.RuleIdType, device_id_list string) {
	//TODO
}

//从Cache中去除指定规则的白名单
func RemoveDeviceIDList(rule_id model.RuleIdType) {
	//TODO
}

//指定的device_id是否在指定规则的白名单列表中
func CheckDeviceIDListList(rule_id model.RuleIdType, device_id string) (hit bool) {
	//TODO
	return false
}
