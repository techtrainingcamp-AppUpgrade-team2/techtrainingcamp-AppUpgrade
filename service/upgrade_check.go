package service

import (
	"app_upgrade/model"
	"strconv"
	"strings"
)

// @title             UpgradeCheck
// @description       检查所有规则，命中则返升级包的信息
// @auth              卢品洲         2021/11/7
// @param             cd            客户端数据
// @return            hit           客户请求是否命中
// @return            pkg           升级包的信息
func UpgradeCheck(cd *model.ClientData) (hit bool, pkg *model.PkgData) {
	rules := GetRulesFromDB() //获取所有规则

	var hit_rule *model.Rule = nil
	max_version := cd.VertionCode
	for i := range rules { //匹配最高版本的规则
		if rules[i].State && checkRule(cd, &rules[i]) && versionLess(max_version, rules[i].UpdateVersionCode) {
			hit_rule = &rules[i]
		}
	}

	if hit_rule != nil {
		return true, hit_rule.GetUpdatePackageInfo()
	} else {
		return false, nil
	}
}

// @title             checkRule
// @description       检查某一条客户端信息是否命中指定规则
// @auth              卢品洲         2021/11/7
// @param             cd            客户端数据
// @param             rule          所指定的规则
// @return            -             是否命中
func checkRule(cd *model.ClientData, rule *model.Rule) bool {
	//匹配app_id、平台、渠道
	if cd.Aid != rule.Aid || cd.DevicePlatform != rule.Platform || cd.Channel != rule.Channel {
		return false
	}
	//设备ID是否在规则的白名单内
	// if !CheckDeviceIDListList(rule.Rid, cd.DeviceId) {
	// 	return false
	// }
	//匹配CPU架构
	if tmp, _ := strconv.Atoi(rule.CpuArch); cd.CpuArch != tmp {
		return false
	}
	//若为安卓平台，则检查安卓版本号是否落在规则所指定的范围[MinOsApi,MaxOsApi]内
	if cd.DevicePlatform == "Android" && (cd.OsApi < rule.MinOsApi || cd.OsApi > rule.MaxOsApi) {
		return false
	}
	//检查客户端当前版本号是否在[MinUpdateVersionCode,MaxUpdateVersionCode]内
	if versionLess(cd.VertionCode, rule.MinUpdateVersionCode) || versionLess(rule.MaxUpdateVersionCode, cd.VertionCode) {
		return false
	}

	return true
}

// @title             versionLess
// @description       判断两个版本号的大小关系
// @auth              卢品洲         2021/11/7
// @param             v1,v2         待比较的两个版本号
// @return            -             v1是否严格小于v2
func versionLess(v1, v2 string) bool {
	//分割版本号并“对齐”
	arr1, arr2 := strings.Split(v1, "."), strings.Split(v2, ".")
	for len(arr1) < len(arr2) {
		arr1 = append(arr1, "0")
	}
	for len(arr2) < len(arr1) {
		arr2 = append(arr2, "0")
	}

	//比较大小
	for i := range arr1 {
		n1, _ := strconv.Atoi(arr1[i])
		n2, _ := strconv.Atoi(arr2[i])
		if n1 < n2 {
			return true
		} else if n1 > n2 {
			return false
		}
	}

	return false
}
