package system

import (
	"Chinese_Learning_App/global"
	"Chinese_Learning_App/model/common/response"
	"Chinese_Learning_App/model/system"
	sysRes "Chinese_Learning_App/model/system/response"
	"Chinese_Learning_App/utils"
	"os"

	"github.com/gin-gonic/gin"
)

type SysAmusementVideoApi struct{}

// AddAmusementVideoApi 添加娱乐视频
func (s *SysAmusementVideoApi) AddAmusementVideoApi(ctx *gin.Context) {
	// todo： 上传娱乐视频
	video, err := ctx.FormFile("video")         // 获取视频资源
	videoIcon, err := ctx.FormFile("videoIcon") // 获取图片资源
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	// 获取其他的参数
	videoName := ctx.PostForm("videoName")
	videoType := ctx.PostForm("type")
	content := ctx.PostForm("content")
	// 开始上传minio到资源
	path := "./video/" + video.Filename
	err = ctx.SaveUploadedFile(video, path)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	// 上传图片资源
	iconUrl, err := utils.UploadToMinio(global.CLA_CONFIG.Minio.BucketName,
		videoName+videoIcon.Filename,
		videoIcon,
		"application/octet-stream")
	// 上传视频资源
	videoUrl, err := utils.UploadVideoToMinio(path,
		global.CLA_CONFIG.Minio.BucketName,
		videoName+video.Filename,
		global.CLA_Minio_Client)

	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
	}
	err = sysAmusementVideoService.AddAmusementVideo(&system.SysAmusementVodeo{
		VideoName:    videoName,
		VideoUrl:     videoUrl,
		VideoIconUrl: iconUrl,
		Type:         videoType,
		Content:      content,
	})
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	err = os.Remove(path)
	if err != nil {
		response.FailWithMessage("上传有问题", ctx)
		return
	}
	response.OkWithMessage("测试成功", ctx)
	return
}

// GetAmusementVideoApi 获取娱乐视频信息
func (s SysAmusementVideoApi) GetAmusementVideoApi(ctx *gin.Context) {
	// todo: 暂时将所有视频查询出来，后面再优化
	list, err := sysAmusementVideoService.GetAmusementVideo()
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	response.OkWithDetailed(sysRes.AmusementVideoListResponse{
		List:  list,
		Total: len(list),
	}, "获取成功", ctx)
	return
}
