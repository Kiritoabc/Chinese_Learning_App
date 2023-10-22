package system

import (
	"Chinese_Learning_App/global"
	"Chinese_Learning_App/model/common/response"
	"Chinese_Learning_App/model/system"
	sysReq "Chinese_Learning_App/model/system/request"
	sysRes "Chinese_Learning_App/model/system/response"
	"Chinese_Learning_App/utils"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

type SysUserApi struct{}

// Register
// @Tags      Register
// @Summary   注册账号
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200   {object}   "创建角色,返回包括系统角色详情"
// @Router    /user/register [post]
func (a *SysUserApi) Register(ctx *gin.Context) {
	// todo: 暂时不考虑使用验证码注册

	var r sysReq.Register
	// 获取参数
	err := ctx.ShouldBindJSON(&r)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	// 参数的校验
	err = utils.Verify(r, utils.RegisterVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	// 注册参数赋值
	user := &system.SysUser{
		Username:  r.Username,
		Password:  r.Password,
		NickName:  r.NickName,
		HeaderImg: r.HeaderImg,
		Enable:    r.Enable,
		Phone:     r.Phone,
		Email:     r.Email,
	}
	// 调用注册服务
	userReturn, err := sysUserService.Register(*user)
	if err != nil {
		global.CLA_LOG.Error("注册失败!", zap.Error(err))
		response.FailWithDetailed(sysRes.SysUserResponse{User: userReturn}, "注册失败", ctx)
		return
	}
	response.OkWithDetailed(sysRes.SysUserResponse{User: userReturn}, "注册成功", ctx)
}
