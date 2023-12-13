package system

import (
	"Chinese_Learning_App/global"

	"gorm.io/gorm"
)

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

func (tv *SysTeachingVideo) AfterSave(tx *gorm.DB) error {
	if tv.ParentId == 0 {
		tx.Model(tx).Update("parent_id", tv.ID).Update("number", 1).Update("episode", 1)
		return nil
	}
	var p *SysTeachingVideo
	// 查询到父
	tx.Model(&SysTeachingVideo{}).First(&p, "id = ?", tv.ParentId)
	// 更新父类的number
	tx.Model(p).Update("number", p.Number+1)
	// 更新子类的number和episode
	tx.Model(tx).Update("number", p.Number+1).Update("episode", p.Number+1)
	// 是否更新所有子？暂时不考虑
	return nil
}
