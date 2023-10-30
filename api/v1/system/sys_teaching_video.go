package system

import (
	"Chinese_Learning_App/global"
	"Chinese_Learning_App/model/common/request"
	"Chinese_Learning_App/model/common/response"
	"Chinese_Learning_App/model/system"
	"Chinese_Learning_App/utils"
	"os"
	"strconv"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

type SysTeachingVideoApi struct{}

// AddTeachingVideoApi
// @Tags AddTeachingVideoApi
// @Summary 新增SysTeachingVideo
// @router /teachingVideo/upload [post]
// @Success 200 "新增成功"
func (s *SysTeachingVideoApi) AddTeachingVideoApi(ctx *gin.Context) {
	// todo: 修改成批量上传
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
	err = sysTeachingVideoService.AddTeachingVideo(*teachingVideo)
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

// SearchTeachingVideoList
// @Tags SysTeachingVideo
// @router /teachingVideo/getTeachingVideoList [post]
// @Summary 查询视频
func (s *SysTeachingVideoApi) SearchTeachingVideoList(ctx *gin.Context) {
	var pageInfo request.PageInfo
	err := ctx.ShouldBindJSON(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	// 参数验证
	err = utils.Verify(pageInfo, utils.PageInfoVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
	}

	teachingVideoList, total, err := sysTeachingVideoService.GetTeachingVideoList(pageInfo)
	if err != nil {
		global.CLA_LOG.Error("获取失败", zap.Error(err))
		response.FailWithMessage("获取失败", ctx)
		return
	}

	response.OkWithDetailed(response.PageResult{
		List:     teachingVideoList,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", ctx)
}

func (s *SysTeachingVideoApi) SearchAllParentVideoList(ctx *gin.Context) {

}
