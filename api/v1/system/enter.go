package system

import "Chinese_Learning_App/service"

type ApiGroup struct {
	SysTeachingVideoApi
	SysUserApi
	SysAmusementVideoApi
}

var (
	sysTeachingVideoService = service.ServiceGroupApp.SystemServiceGroup.TeachingVideoService
	sysUserService          = service.ServiceGroupApp.SystemServiceGroup.UserService
)
