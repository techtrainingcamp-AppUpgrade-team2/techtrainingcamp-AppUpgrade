package model

type DeviceIdType string

type ClientData struct {
	Version           string       `form:"version"`             //api版本
	DevicePlatform    string       `form:"device_platform"`     //设备平台
	DeviceId          DeviceIdType `form:"device_id"`           //设备ID，一般是16位的由数字和字母构成的字符串
	OsApi             int          `form:"os_api"`              //安卓系统版本
	Channel           string       `form:"channel"`             //渠道
	VertionCode       string       `form:"vertion_code"`        //应用大版本，如8.1.4
	UpdateVersionCode string       `form:"update_version_code"` //应用小版本，如8.1.4.01
	Aid               int          `form:"aid"`                 //app的唯一标识
	CpuArch           int          `form:"cpu_arch"`            //设备的CPU结构，如32/64
}
