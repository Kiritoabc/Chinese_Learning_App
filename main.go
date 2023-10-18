package main

import (
	"Chinese_Learning_App/core"
	"Chinese_Learning_App/global"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// @title				swagger API
// @Version				1.1.1
// @description			暂无
// @BasePath			/
func main() {
	global.CLA_VP = core.Viper() // 初始化Viper
	global.CLA_LOG = core.Zap()  // 初始化Zap日志库
	S := gin.Default()
	// 默认不拦截，
	//todo：暂时不对跨域做配置
	S.Use(cors.Default())
	S.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "服务启动成功"})
	})
	err := S.Run(":8080")
	if err != nil {
		fmt.Println("服务器启动失败！")
	}
}
