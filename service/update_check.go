package service

import (
	"app_upgrade/model"
	"strconv"
	"strings"
)

// @title             UpdateCheck
// @description       检查所有规则，命中则返升级包的信息
// @auth              卢品洲         2021/11/7  10:30
// @param             cd            model.ClientData    客户端数据
// @return            hit           bool                客户请求是否命中
// @return            pkg           model.PkgData       升级包的信息
func UpdateCheck(cd *model.ClientData) (hit bool, pkg *model.PkgData) {
	rules := GetRulesFromDB() //获取所有规则，TODO：性能？

	for _, rule := range rules { //遍历规则列表
		hit = checkRule(cd, rule) //若某条规则与客户端信息匹配，则返回升级包的信息
		if hit {
			return true, &model.PkgData{
				DownloadUrl:       rule.DownloadUrl,
				UpdateVersionCode: rule.UpdateVersionCode,
				Md5:               rule.Md5,
				Title:             rule.Title,
				UpdateTips:        rule.UpdateTips,
			}
		}
	}
	return false, nil
}

// @title             checkRule
// @description       检查某一条客户端信息是否命中指定规则
// @auth              卢品洲         2021/11/7
// @param             cd            model.ClientData    客户端数据
// @param             rule          model.Rule          所指定的规则
// @return            -             bool                是否命中
func checkRule(cd *model.ClientData, rule *model.Rule) bool {
	//匹配app_id、平台、渠道
	if cd.Aid != rule.Aid || cd.DevicePlatform != rule.Platform || cd.Channel != rule.Channel {
		return false
	}
	//设备ID是否在规则的白名单内
	if !CheckDeviceIDList(rule.Rid, string(cd.DeviceId)) {
		return false
	}
	//匹配CPU架构
	if tmp, _ := strconv.Atoi(rule.CpuArch); cd.CpuArch != tmp {
		return false
	}
	//若为安卓平台，则检查安卓版本号是否落在规则所指定的范围[MinOsApi,MaxOsApi]内
	if cd.DevicePlatform == "Android" && (cd.OsApi < rule.MinOsApi || cd.OsApi > rule.MaxOsApi) {
		return false
	}
	//检查客户端当前版本号是否在[MinUpdateVersionCode,MaxUpdateVersionCode]内
	if version_less(cd.VertionCode, rule.MinUpdateVersionCode) || version_less(rule.MaxUpdateVersionCode, cd.VertionCode) {
		return false
	}

	return true
}

// @title             version_less
// @description       判断两个版本号的大小关系
// @auth              卢品洲         2021/11/7
// @param             v1,v2         string    待比较的两个版本号,如8.1.3、9.4.6.1
// @return            -             bool      v1是否严格小于v2
func version_less(v1, v2 string) bool {
	//分割版本号并“对齐”
	arr1, arr2 := strings.Split(v1, "."), strings.Split(v2, ".")
	for len(arr1) < len(arr2) {
		arr1 = append(arr1, "0")
	}
	for len(arr2) < len(arr1) {
		arr2 = append(arr2, "0")
	}

	for i, _ := range arr1 {
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
