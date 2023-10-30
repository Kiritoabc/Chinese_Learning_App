package system

import "Chinese_Learning_App/global"

type SysTeachingVideo struct {
	global.GVA_MODEL
	VideoName    string `json:"videoName" gorm:"comment:视频名称"`
	VideoUrl     string `json:"videoUrl" gorm:"comment:视频地址"`
	VideoIconUrl string `json:"videoIcon" gorm:"comment:视频图标地址"`
	Group        string `json:"group" gorm:"comment:视频分组"`
	Type         int    `json:"type" gorm:"comment:类型"`
	Number       int    `json:"number" gorm:"comment:集数"`
	Episode      int    `json:"episode" gorm:"comment:第几集"`
	ParentId     uint   `json:"parentId" gorm:"comment:父Id"`
	Content      string `json:"content" gorm:"comment:内容"`
}

func (SysTeachingVideo) TableName() string {
	return "sys_teaching_videos"
}
