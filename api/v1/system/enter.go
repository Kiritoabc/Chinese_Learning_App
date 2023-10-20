package system

import "Chinese_Learning_App/service"

type ApiGroup struct {
	SysTeachingVideoApi
}

var (
	sysTeachingVideoService = service.ServiceGroupApp.SystemServiceGroup.TeachingVideoService
)
