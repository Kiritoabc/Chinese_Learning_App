package system

import "Chinese_Learning_App/global"

type SysAmusementVodeo struct {
	global.GVA_MODEL
	VideoName    string `json:"videoName" gorm:"comment: 娱乐视频名称"`
	VideoUrl     string `json:"videoUrl" gorm:"comment: 视频地址"`
	VideoIconUrl string `json:"videoIconUrl" gorm:"comment: 封面地址"`
	Type         string `json:"type" gorm:"comment: 视频类型"`
	Content      string `json:"content" gorm:"comment:视频的内容"`
}
