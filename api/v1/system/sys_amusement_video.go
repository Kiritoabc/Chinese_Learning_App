package system

import (
	"Chinese_Learning_App/model/common/response"

	"github.com/gin-gonic/gin"
)

type SysAmusementVideoApi struct{}

// AddAmusementVideoApi 添加娱乐视频
func (s *SysAmusementVideoApi) AddAmusementVideoApi(ctx *gin.Context) {
	response.OkWithMessage("测试成功", ctx)
	return
}

// GetAmusementVideoApi 获取娱乐视频信息
func (s SysAmusementVideoApi) GetAmusementVideoApi(ctx *gin.Context) {
	response.OkWithMessage("测试成功", ctx)
	return
}
