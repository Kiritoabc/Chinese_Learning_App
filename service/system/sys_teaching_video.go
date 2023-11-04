package system

import (
	"Chinese_Learning_App/global"
	"Chinese_Learning_App/model/common/request"
	"Chinese_Learning_App/model/system"
)

type TeachingVideoService struct {
}

// AddTeachingVideo
// 添加教学视频
func (t *TeachingVideoService) AddTeachingVideo(teachingVideo system.SysTeachingVideo) error {

	// gorm 保存到数据库
	global.CLA_DB.Save(&teachingVideo)
	return nil
}

// GetTeachingVideoList
// 分页获取数据
func (t *TeachingVideoService) GetTeachingVideoList(info request.PageInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.CLA_DB.Model(&system.SysTeachingVideo{})
	if err = db.Where("type = ?", "0").Count(&total).Error; total == 0 || err != nil {
		return
	}
	var teachingVideo []system.SysTeachingVideo
	err = db.Limit(limit).Offset(offset).Find(&teachingVideo).Error
	if err != nil {
		return
	}
	return teachingVideo, total, err
}

func (t *TeachingVideoService) SearchAllParentVideoList() (list []system.SysTeachingVideo, err error) {
	db := global.CLA_DB.Model(&system.SysTeachingVideo{})
	// 查询所有的第一集的信息
	err = db.Where("episode = ?", "1").Find(&list).Error
	return
}

func (t *TeachingVideoService) SearchAllSonVideoList(ParentId int) (list []system.SysTeachingVideo, total int64, err error) {
	db := global.CLA_DB.Model(&system.SysTeachingVideo{})
	if err = db.Where("parent_id = ?", ParentId).Find(&list).Count(&total).Error; total == 0 || err != nil {
		return
	}

	return list, total, err
}
