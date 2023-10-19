package initialize

import (
	"Chinese_Learning_App/global"
	"Chinese_Learning_App/router"
	"net/http"

	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/swag/example/basic/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// 初始化总路由

func Routers() *gin.Engine {
	Router := gin.Default()
	// todo: 路由暂时不做跨域拦截配置，全部放行
	Router.Use(cors.Default())
	//InstallPlugin(Router) // 安装插件
	systemRouter := router.RouterGroupApp.System
	//exampleRouter := router.RouterGroupApp.Example
	// swagger 配置
	docs.SwaggerInfo.BasePath = global.CLA_CONFIG.System.RouterPrefix
	Router.GET(global.CLA_CONFIG.System.RouterPrefix+"/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	global.CLA_LOG.Info("register swagger handler")
	//
	PublicGroup := Router.Group(global.CLA_CONFIG.System.RouterPrefix)
	{
		// 健康监测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, "ok")
		})
	}
	{
		systemRouter.InitBaseRouter(PublicGroup) // 注册基础功能路由 不做鉴权
		systemRouter.InitInitRouter(PublicGroup) // 自动初始化相关
	}
	//
	PrivateGroup := Router.Group(global.CLA_CONFIG.System.RouterPrefix)
	{
		PrivateGroup.GET("/hello", func(c *gin.Context) {
			c.JSON(http.StatusOK, "hello world")
		})
	}
	global.CLA_LOG.Info("router register success")
	return Router
}
