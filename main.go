package main

import (
	"Chinese_Learning_App/core"
	"Chinese_Learning_App/global"
	"Chinese_Learning_App/initialize"

	"go.uber.org/zap"
)

// @title				swagger API
// @Version				1.1.1
// @description			暂无
// @BasePath			/
func main() {
	global.CLA_VP = core.Viper() // 初始化Viper
	global.CLA_LOG = core.Zap()  // 初始化Zap日志库
	zap.ReplaceGlobals(global.CLA_LOG)
	global.GVA_DB = initialize.Gorm() // gorm连接数据库
	initialize.DBList()
	if global.GVA_DB != nil {
		initialize.RegisterTables() // 初始化表
		// 程序结束前关闭数据库链接
		db, _ := global.GVA_DB.DB()
		defer db.Close()
	}
	core.RunWindowsServer()
}
