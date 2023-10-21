package system

import (
	"Chinese_Learning_App/global"
	"Chinese_Learning_App/model/common/response"
	"Chinese_Learning_App/model/system"
	"Chinese_Learning_App/utils"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SysTeachingVideoApi struct{}

// AddTeachingVideoApi
// @Tags AddTeachingVideoApi
// @Summary 新增SysTeachingVideo
// @router /teachingVideo/upload [post]
// @Success 200 "新增成功"
func (s *SysTeachingVideoApi) AddTeachingVideoApi(ctx *gin.Context) {
	// 注意这里只测试单个视频
	video, _ := ctx.FormFile("video")
	videoIcon, _ := ctx.FormFile("videoIcon")
	videoName := ctx.PostForm("videoName")
	group := ctx.PostForm("group")
	var videoType int
	videoType, _ = strconv.Atoi(ctx.PostForm("videoType"))
	teachingVideo := &system.SysTeachingVideo{
		VideoName: videoName,
		Group:     group,
		Type:      videoType,
	}
	path := "./video/" + video.Filename
	err := ctx.SaveUploadedFile(video, path)
	if err != nil {
		response.FailWithMessage("保存视频失败", ctx)
		return
	}

	// 上传到minio
	iconUrl, err := utils.UploadToMinio("test", teachingVideo.VideoName+videoIcon.Filename, videoIcon, "application/octet-stream")
	if err != nil {
		return
	}
	videoUrl, err := utils.UploadVideoToMinio(path, "test", teachingVideo.VideoName+video.Filename, global.CLA_Minio_Client)
	teachingVideo.VideoIconUrl = iconUrl
	teachingVideo.VideoUrl = videoUrl

	// 调用 sysTeachingVideoService 保存到数据库
	err = sysTeachingVideoService.AddTeachingVideo(teachingVideo)
	if err != nil {
		response.FailWithMessage("上传视频失败", ctx)
		return
	}
	err = os.Remove(path)
	if err != nil {
		response.FailWithMessage("上传有问题", ctx)
		return
	}

	response.OkWithMessage("上传视频成功", ctx)
}
