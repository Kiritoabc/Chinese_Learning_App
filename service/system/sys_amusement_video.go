package system

import (
	"Chinese_Learning_App/global"
	"Chinese_Learning_App/model/system"
)

type AmusementVideoService struct {
}

// AddAmusementVideo 上传视频
func (s *AmusementVideoService) AddAmusementVideo(video *system.SysAmusementVodeo) (err error) {
	global.CLA_DB.Model(&system.SysAmusementVodeo{}).Save(&video)
	return
}

// GetAmusementVideo 获取视频信息
func (s AmusementVideoService) GetAmusementVideo() (list []system.SysAmusementVodeo, err error) {
	db := global.CLA_DB.Model(&system.SysAmusementVodeo{})
	db.Find(&list)
	return
}
