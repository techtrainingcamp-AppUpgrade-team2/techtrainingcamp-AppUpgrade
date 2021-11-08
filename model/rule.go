package model

type RuleIdType int

type Rule struct {
	//标识信息
	State bool       `json:"state"` //规则是否可用
	Rid   RuleIdType `json:"rid"`   //规则唯一标识

	//包推送要求
	Aid      int    `json:"aid"`      //app唯一标识
	Platform string `json:"platjson"` //平台，Android/iOS
	Channel  string `json:"channel"`  //渠道号

	DeviceIdList         string `json:"device_id_list"`          //白名单
	MinUpdateVersionCode string `json:"min_update_version_code"` //旧app版本区间
	MaxUpdateVersionCode string `json:"max_update_version_code"`
	MinOsApi             int    `json:"min_os_api"` //支持的操作系统版本区间（针对Android）
	MaxOsApi             int    `json:"max_os_api"`
	CpuArch              string `json:"cpu_arch"` //CPU架构

	//包信息，TODO:以下字段是否冗余？
	DownloadUrl       string `json:"download_url"`        //下载链接
	UpdateVersionCode string `json:"update_version_code"` //版本号
	Md5               string `json:"md5"`                 //哈希值
	Title             string `json:"title"`               //弹窗标题
	UpdateTips        string `json:"update_tips"`         //弹窗文本
}
