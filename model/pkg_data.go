package model

type PkgData struct {
	DownloadUrl       string `form:"download_url"`        //下载链接
	UpdateVersionCode string `form:"update_version_code"` //版本号
	Md5               string `form:"md5"`                 //哈希值
	Title             string `form:"title"`               //弹窗标题
	UpdateTips        string `form:"update_tips"`         //弹窗文本
}
