package initialize

import (
	"Chinese_Learning_App/global"
	"Chinese_Learning_App/model/system"
	"os"

	"go.uber.org/zap"

	"gorm.io/gorm"
)

// Gorm 初始化数据库并产生数据库全局变量
// Author SliverHorn
func Gorm() *gorm.DB {
	switch global.CLA_CONFIG.System.DbType {
	case "mysql":
		return GormMysql()
	//case "pgsql":
	//	return GormPgSql()
	//case "oracle":
	//	return GormOracle()
	//case "mssql":
	//	return GormMssql()
	//case "sqlite":
	//	return GormSqlite()
	default:
		return GormMysql()
	}
}

// RegisterTables 注册数据库表专用
func RegisterTables() {
	db := global.CLA_DB
	err := db.AutoMigrate(
		// 系统模块表
		system.SysTeachingVideo{},
		system.SysUser{},
		system.SysAmusementVodeo{},
	)
	if err != nil {
		global.CLA_LOG.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	global.CLA_LOG.Info("register table success")
}
