package model

type PkgData struct {
	DownloadUrl       string `json:"download_url"`        //下载链接
	UpdateVersionCode string `json:"update_version_code"` //版本号
	Md5               string `json:"md5"`                 //哈希值
	Title             string `json:"title"`               //弹窗标题
	UpdateTips        string `json:"update_tips"`         //弹窗文本
}
