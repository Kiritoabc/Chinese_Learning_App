package system

import (
	"Chinese_Learning_App/global"
	"Chinese_Learning_App/model/common/request"
	"Chinese_Learning_App/model/common/response"
	"Chinese_Learning_App/model/system"
	sysRes "Chinese_Learning_App/model/system/response"
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
// todo: 上传不是重点，暂时不做要求
func (s *SysTeachingVideoApi) AddTeachingVideoApi(ctx *gin.Context) {
	// todo: 修改成批量上传
	// 注意这里只测试单个视频
	video, _ := ctx.FormFile("video")
	videoIcon, _ := ctx.FormFile("videoIcon")
	videoName := ctx.PostForm("videoName")
	group := ctx.PostForm("group")
	parentId, _ := strconv.Atoi(ctx.PostForm("parentId"))
	var videoType int
	videoType, _ = strconv.Atoi(ctx.PostForm("videoType"))
	var teachingVideo = &system.SysTeachingVideo{
		ParentId:  uint(parentId),
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
// @Tags SearchTeachingVideoList
// @router /teachingVideo/getTeachingVideoList [post]
// @Summary 查询视频？好像没用了哦
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

// SearchAllParentVideoList
// @Tags SearchAllParentVideoList
// @router /teachingVideo/getTeachingParentVideoList [post]
// @Summary 查询所有父视频
func (s *SysTeachingVideoApi) SearchAllParentVideoList(ctx *gin.Context) {
	var tempList []sysRes.TeachingVideoInfoResponse

	list, err := sysTeachingVideoService.SearchAllParentVideoList()

	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	total := 0
	// 将list转换成tempList
	for _, item := range list {
		tempItem := sysRes.TeachingVideoInfoResponse{
			ID:           item.ID,
			VideoName:    item.VideoName,
			Number:       item.Number,
			Content:      item.Content,
			VideoIconUrl: item.VideoIconUrl,
		}
		total++
		tempList = append(tempList, tempItem)
	}
	response.OkWithDetailed(sysRes.TeachingVideoInfoListResponse{
		List:  tempList,
		Total: total,
	}, "获取成功", ctx)
}

// SearchAllSonVideoList
// @Tags SearchAllSonVideoList
// @router /teachingVideo/getTeachingSonVideoList [post]
// @Summary 查询所有子视频
func (s *SysTeachingVideoApi) SearchAllSonVideoList(ctx *gin.Context) {
	ParentId, err := strconv.Atoi(ctx.PostForm("ParentId"))
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	list, total, err := sysTeachingVideoService.SearchAllSonVideoList(ParentId)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}

	response.OkWithDetailed(sysRes.TeachingVideoListResponse{
		List:  list,
		Total: total,
	}, "获取成功", ctx)
}
