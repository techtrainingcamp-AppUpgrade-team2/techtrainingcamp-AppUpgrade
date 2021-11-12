package service

import (
	"app_upgrade/model"
	"testing"
)

func TestVersionLess(t *testing.T) {
	cases := []struct {
		Name     string
		V1, V2   string
		Expected bool
	}{
		{"1", "8.1.4", "8.1.3", false},
		{"2", "8.1.1", "8.1.3", true},
		{"3", "8.1.2", "8.1.1.11", false},
		{"4", "1.1.4", "2.1.3", true},
		{"5", "1.1.4", "1.1.4.2", true},
		{"6", "1.4", "1.1.4.2", false},
		{"7", "2.1", "1.1.3.8", false},
		{"8", "3", "2.4.6", false},
		{"9", "4", "5.7.6", true},
	}

	for _, c := range cases {
		if ret := versionLess(c.V1, c.V2); ret != c.Expected {
			t.Error(c.Name, c.V1, c.V2)
		}
	}
}

//测试除白名单外的规则匹配逻辑是否正确
func TestCheckRule(t *testing.T) {
	cases := []struct {
		no       int
		cd       *model.ClientData
		rule     *model.Rule
		expected bool
	}{
		{
			no: 1,
			cd: &model.ClientData{
				Version:           "1.0",
				DevicePlatform:    "Android",
				DeviceId:          model.DeviceIdType("abcdefghijklmnopq"),
				OsApi:             11,
				Channel:           "xiaomi",
				VertionCode:       "1.1.1",
				UpdateVersionCode: "",
				Aid:               1,
				CpuArch:           64,
			},
			rule: &model.Rule{
				State:    true,
				Rid:      model.RuleIdType(1),
				Aid:      1,
				Platform: "Android",
				Channel:  "meizu",

				DeviceIdList:         "abcdefghijklmnopq,abcdefg",
				MinUpdateVersionCode: "1.1",
				MaxUpdateVersionCode: "1.3",
				MinOsApi:             9,
				MaxOsApi:             12,
				CpuArch:              "64",

				DownloadUrl:       "www.baidu.com",
				UpdateVersionCode: "1.1.3",
				Md5:               "md5",
				Title:             "版本更新",
				UpdateTips:        "新版本发布啦，快来更新吧！",
			},
			expected: false,
		},
		{
			no: 2,
			cd: &model.ClientData{
				Version:           "1.0",
				DevicePlatform:    "iOS",
				DeviceId:          model.DeviceIdType("abcdefghijklmnopq"),
				OsApi:             -1,
				Channel:           "apple",
				VertionCode:       "1.1.1",
				UpdateVersionCode: "",
				Aid:               1,
				CpuArch:           64,
			},
			rule: &model.Rule{
				State:    true,
				Rid:      model.RuleIdType(1),
				Aid:      1,
				Platform: "Android",
				Channel:  "meizu",

				DeviceIdList:         "abcdefghijklmnopq,abcdefg",
				MinUpdateVersionCode: "1.1",
				MaxUpdateVersionCode: "1.3",
				MinOsApi:             9,
				MaxOsApi:             12,
				CpuArch:              "64",

				DownloadUrl:       "www.baidu.com",
				UpdateVersionCode: "1.1.3",
				Md5:               "md5",
				Title:             "版本更新",
				UpdateTips:        "新版本发布啦，快来更新吧！",
			},
			expected: false,
		},
		{
			no: 3,
			cd: &model.ClientData{
				Version:           "1.0",
				DevicePlatform:    "Android",
				DeviceId:          model.DeviceIdType("abcdefghijklmnopq"),
				OsApi:             11,
				Channel:           "xiaomi",
				VertionCode:       "1.1.1",
				UpdateVersionCode: "",
				Aid:               1,
				CpuArch:           64,
			},
			rule: &model.Rule{
				State:    true,
				Rid:      model.RuleIdType(1),
				Aid:      1,
				Platform: "Android",
				Channel:  "xiaomi",

				DeviceIdList:         "abcdefghijklmnopq,abcdefg",
				MinUpdateVersionCode: "1.1.2",
				MaxUpdateVersionCode: "1.2.4",
				MinOsApi:             9,
				MaxOsApi:             11,
				CpuArch:              "64",

				DownloadUrl:       "www.baidu.com",
				UpdateVersionCode: "1.1.3",
				Md5:               "md5",
				Title:             "版本更新",
				UpdateTips:        "新版本发布啦，快来更新吧！",
			},
			expected: false,
		},
		//...
	}

	for _, c := range cases {
		if ret := checkRule(c.cd, c.rule); ret != c.expected {
			t.Errorf("第%d组用例不通过", c.no)
		}
	}
}
