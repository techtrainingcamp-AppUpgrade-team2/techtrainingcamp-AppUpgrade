package model

//PkgData 升级包的信息
type PkgData struct {
	DownloadUrl       string `form:"download_url" json:"download_url"`               //下载链接
	UpdateVersionCode string `form:"update_version_code" json:"update_version_code"` //版本号
	Md5               string `form:"md5" json:"md5"`                                 //哈希值
	Title             string `form:"title" json:"title"`                             //弹窗标题
	UpdateTips        string `form:"update_tips" json:"update_tips"`                 //弹窗文本

}
