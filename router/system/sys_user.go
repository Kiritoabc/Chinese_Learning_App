package system

import (
	v1 "Chinese_Learning_App/api/v1"

	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (s *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userGroup := Router.Group("/user")
	userApi := v1.ApiGroupApp.SystemApiGroup.SysUserApi

	{
		userGroup.POST("/register", userApi.Register) // 注册账号
		userGroup.POST("/login", userApi.Login)       // 用户登录
	}
}
