package system

import (
	"Chinese_Learning_App/global"
	"Chinese_Learning_App/model/common/response"
	"Chinese_Learning_App/model/system"
	sysReq "Chinese_Learning_App/model/system/request"
	sysRes "Chinese_Learning_App/model/system/response"

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
	// 参数的校验 todo: 暂时不验证参数
	//err = utils.Verify(r, utils.RegisterVerify)
	//if err != nil {
	//	response.FailWithMessage(err.Error(), ctx)
	//	return
	//}
	// 注册参数赋值
	user := &system.SysUser{
		Username:  r.Username,
		Password:  r.Password,
		NickName:  r.NickName,
		HeaderImg: "https://ts1.cn.mm.bing.net/th/id/R-C.e3506e1ab5441c8c9c43facfba9ff1ab?rik=JKbGTz1jQYhGuA&riu=http%3a%2f%2fwww.gx8899.com%2fuploads%2fallimg%2f2018031109%2fqzfrr5ly3af.jpg&ehk=izdYHSNuzDYZQF2HdyA6xZ3Jgz5lXaj%2faZHquTh2FvU%3d&risl=&pid=ImgRaw&r=0&sres=1&sresct=1",
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

func (a *SysUserApi) Login(ctx *gin.Context) {
	var r sysReq.Register
	// 获取参数
	err := ctx.ShouldBindJSON(&r)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	// 注册参数赋值
	user := &system.SysUser{
		Username: r.Username,
		Password: r.Password,
	}

	userReturn, err := sysUserService.Login(*user)
	if err != nil {
		global.CLA_LOG.Error("登录失败!", zap.Error(err))
		response.FailWithDetailed(sysRes.SysUserResponse{
			User: userReturn,
		}, "登陆失败", ctx)
		return
	}
	response.OkWithDetailed(sysRes.SysUserResponse{User: userReturn}, "登录成功", ctx)
}
