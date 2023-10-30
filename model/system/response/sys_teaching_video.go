package response

import "Chinese_Learning_App/model/system"

// 用于返回所有父类视频的集数

type TeachingVideoInfoResponse struct {
	VideoName    string `json:"videoName"`    //视频名称
	VideoIconUrl string `json:"videoIconUrl"` //视频的封面URL
	Number       int    `json:"number"`       //总集数
	Content      string `json:"content"`      //视频描述的内容
	ID           uint   `json:"ID"`           //视频的ID
}

type TeachingVideoInfoListResponse struct {
	List  []TeachingVideoInfoResponse `json:"list"`
	Total int                         `json:"total"`
}

type TeachingVideoListResponse struct {
	List  []system.SysTeachingVideo `json:"list"`
	Total int64                     `json:"total"`
}
