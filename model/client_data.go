package model

type DeviceIdType string

type ClientData struct {
	Version           string       `json:"version"`             //api版本
	DevicePlatform    string       `json:"device_platjson"`     //设备平台
	DeviceId          DeviceIdType `json:"device_id"`           //设备ID，一般是16位的由数字和字母构成的字符串
	OsApi             int          `json:"os_api"`              //安卓系统版本
	Channel           string       `json:"channel"`             //渠道
	VertionCode       string       `json:"vertion_code"`        //应用大版本，如8.1.4
	UpdateVersionCode string       `json:"update_version_code"` //应用小版本，如8.1.4.01
	Aid               int          `json:"aid"`                 //app的唯一标识
	CpuArch           int          `json:"cpu_arch"`            //设备的CPU结构，如32/64
}
