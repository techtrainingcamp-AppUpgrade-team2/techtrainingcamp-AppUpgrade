package service

import (
	"app_upgrade/model"
	"strings"
)

//往Cache中添加指定规则的白名单
func AddDeviceIDList(rule_id model.RuleIdType, device_id_list string) {
	//TODO
	deviceIDCache := model.GetInstance().DeviceIDCache
	device_id_list_slice := strings.Split(device_id_list, ";")
	device_id_list_map := make(map[model.DeviceIdType]struct{})

	for _, device_id := range device_id_list_slice {
		device_id_list_map[model.DeviceIdType(device_id)] = struct{}{}
	}

	deviceIDCache[rule_id] = device_id_list_map
}

//从Cache中去除指定规则的白名单
func RemoveDeviceIDList(rule_id model.RuleIdType) {
	deviceIDCache := model.GetInstance().DeviceIDCache
	delete(deviceIDCache, rule_id)
}

//指定的device_id是否在指定规则的白名单列表中
func CheckDeviceIDListList(rule_id model.RuleIdType, device_id model.DeviceIdType) (hit bool) {
	deviceIDCache := model.GetInstance().DeviceIDCache
	if _, ok := deviceIDCache[rule_id]; !ok {
		rule := GetRuleFromDB(rule_id)
		AddDeviceIDList(rule.Rid, rule.DeviceIdList)
	}
	_, hit = deviceIDCache[rule_id][device_id]
	return hit
}
