package model

type RuleIdType int

type Rule struct {
	//标识信息
	State bool       `form:"state"` //规则是否可用
	Rid   RuleIdType `form:"rid"`   //规则唯一标识

	//包推送要求
	Aid      int    `form:"aid"`      //app唯一标识
	Platform string `form:"platform"` //平台，Android/iOS
	Channel  string `form:"channel"`  //渠道号

	DeviceIdList         string `form:"device_id_list"`          //白名单
	MinUpdateVersionCode string `form:"min_update_version_code"` //旧app版本区间
	MaxUpdateVersionCode string `form:"max_update_version_code"`
	MinOsApi             int    `form:"min_os_api"` //支持的操作系统版本区间（针对Android）
	MaxOsApi             int    `form:"max_os_api"`
	CpuArch              string `form:"cpu_arch"` //CPU架构

	//包信息，TODO:以下字段是否冗余？
	DownloadUrl       string `form:"download_url"`        //下载链接
	UpdateVersionCode string `form:"update_version_code"` //版本号
	Md5               string `form:"md5"`                 //哈希值
	Title             string `form:"title"`               //弹窗标题
	UpdateTips        string `form:"update_tips"`         //弹窗文本
}
