package response

import "Chinese_Learning_App/model/system"

// 用于返回所有父类视频的集数
type AmusementVideoListResponse struct {
	List  []system.SysAmusementVodeo `json:"list"`
	Total int                        `json:"total"`
}
