package service

import (
	"app_upgrade/model"
	"strings"
)

//TODO:选型
var DeviceIDCache map[model.RuleIdType]map[model.DeviceIdType]struct{}

//往Cache中添加指定规则的白名单
func AddDeviceIDList(rule_id model.RuleIdType, device_id_list string) {
	DeviceIDList := make(map[model.DeviceIdType]struct{})
	var Exists = struct{}{}
	sep := ";"
	arr := strings.Split(device_id_list, sep)
	for i := 0; i < len(arr); i++ {
		//print(arr[i])
		var diviceId model.DeviceIdType
		diviceId = model.DeviceIdType(arr[i])
		DeviceIDList[diviceId] = Exists
	}
	if DeviceIDCache == nil {
		DeviceIDCache = make(map[model.RuleIdType]map[model.DeviceIdType]struct{})
	}
	DeviceIDCache[rule_id] = DeviceIDList
	//TODO
}

//从Cache中去除指定规则的白名单
func RemoveDeviceIDList(rule_id model.RuleIdType) {
	delete(DeviceIDCache, rule_id)
	//TODO
}

//指定的device_id是否在指定规则的白名单列表中
func CheckDeviceIDListList(rule_id model.RuleIdType, device_id string) (hit bool) {
	//TODO
	_, ok := DeviceIDCache[rule_id][model.DeviceIdType(device_id)]
	return ok
}
