package service

import (
	"app_upgrade/model"
	"strings"
)

// @title             AddDeviceIDList
// @description       向DeviceIDCache中添加指定规则的白名单
// @auth              高宏宇                2021/11/10
// @param             rule_id              指定规则编号
// @param             device_id_list       指定规则对应的设备白名单
func AddDeviceIDList(rule_id model.RuleIdType, device_id_list string) {
	deviceIDCache := model.GetInstance().DeviceIDCache
	device_id_list_slice := strings.Split(device_id_list, ";")
	device_id_list_map := make(map[model.DeviceIdType]struct{})

	// TODO  构建设备白名单集合
	for _, device_id := range device_id_list_slice {
		device_id_list_map[model.DeviceIdType(device_id)] = struct{}{}
	}

	deviceIDCache[rule_id] = device_id_list_map
}

// @title             RemoveDeviceIDList
// @description       从DeviceIDCache中去除指定规则的白名单
// @auth              高宏宇         2021/11/10
// @param             rule_id       指定规则编号
func RemoveDeviceIDList(rule_id model.RuleIdType) {
	deviceIDCache := model.GetInstance().DeviceIDCache
	delete(deviceIDCache, rule_id)
}

// @title             CheckDeviceIDListList
// @description       检查指定的device_id是否在指定规则的白名单列表中
// @auth              高宏宇           2021/11/10
// @param             rule_id         指定规则编号
// @param             device_id       指定设备编号
// @return            hit             指定的device_id是否在指定规则的设备白名单列表中
func CheckDeviceIDListList(rule_id model.RuleIdType, device_id model.DeviceIdType) (hit bool) {
	deviceIDCache := model.GetInstance().DeviceIDCache
	// TODO  检查指定规则对应的设备白名单是否处在缓存中，不在则载入缓存
	if _, ok := deviceIDCache[rule_id]; !ok {
		rule := GetRuleFromDB(rule_id)
		AddDeviceIDList(rule.Rid, rule.DeviceIdList)
	}
	_, hit = deviceIDCache[rule_id][device_id]
	return hit
}
