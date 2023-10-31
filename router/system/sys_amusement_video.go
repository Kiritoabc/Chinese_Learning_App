package system

import (
	v1 "Chinese_Learning_App/api/v1"

	"github.com/gin-gonic/gin"
)

type AmusementVideoRouter struct{}

func (s *AmusementVideoRouter) InitAmusementVideoRouter(Router *gin.RouterGroup) {
	amusementVideoRouter := Router.Group("amusement")
	amusementVideoApi := v1.ApiGroupApp.SystemApiGroup.SysAmusementVideoApi
	{
		amusementVideoRouter.GET("/get", amusementVideoApi.GetAmusementVideoApi)     // 获取娱乐视频
		amusementVideoRouter.POST("/upload", amusementVideoApi.AddAmusementVideoApi) // 添加娱乐视频信息
	}
}
