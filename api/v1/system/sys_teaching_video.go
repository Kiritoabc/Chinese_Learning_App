package system

import (
	"Chinese_Learning_App/global"
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
	println(videoIcon, teachingVideo)
	global.CLA_LOG.Info(video.Filename)
	err := utils.UploadToMinio("test", videoIcon.Filename, videoIcon, "application/octet-stream")
	path := "./video/" + video.Filename
	err = ctx.SaveUploadedFile(video, path)
	if err != nil {
		return
	}
	// 上传视频
	err = utils.UploadVideoToMinio(path, "test", teachingVideo.VideoName+video.Filename, global.CLA_Minio_Client)
	if err != nil {
		return
	}
	err = os.Remove(path)
	if err != nil {
		return
	}
	if err != nil {
		return
	}
}
