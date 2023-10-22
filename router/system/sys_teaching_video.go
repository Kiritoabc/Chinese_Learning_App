package system

import (
	v1 "Chinese_Learning_App/api/v1"

	"github.com/gin-gonic/gin"
)

type TeachingVideoRouter struct {
}

func (s *TeachingVideoRouter) InitTeachingVideoRouter(Router *gin.RouterGroup) {
	teachingVideoRouter := Router.Group("teachingVideo")

	teachingVideoRouterApi := v1.ApiGroupApp.SystemApiGroup.SysTeachingVideoApi
	{
		teachingVideoRouter.POST("/upload", teachingVideoRouterApi.AddTeachingVideoApi)                   //	上传视频Api
		teachingVideoRouter.POST("/getTeachingVideoList", teachingVideoRouterApi.SearchTeachingVideoList) // 获取教学list的Api
	}
}
