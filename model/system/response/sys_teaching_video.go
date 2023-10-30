package response

// 用于返回所有父类视频的集数

type SysTeachingVideoResponse struct {
	VideoName    string `json:"videoName"`    //视频名称
	VideoIconUrl string `json:"videoIconUrl"` //视频的封面URL
	Number       int64  `json:"number"`       //总集数
	Content      string `json:"content"`      //视频描述的内容
	ID           string `json:"ID"`           //视频的ID
}
