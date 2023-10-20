package global

import (
	"Chinese_Learning_App/config"
	"sync"

	"github.com/minio/minio-go/v7"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// 定义一些全局变量
var (
	CLA_Minio_Client *minio.Client
	CLA_DB           *gorm.DB
	CLA_DBList       map[string]*gorm.DB
	CLA_VP           *viper.Viper  //
	CLA_CONFIG       config.Config // 全局的配置文件
	CLA_LOG          *zap.Logger
	lock             sync.RWMutex
)

// GetGlobalDBByDBName 通过名称获取db list中的db
func GetGlobalDBByDBName(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	return CLA_DBList[dbname]
}

// MustGetGlobalDBByDBName 通过名称获取db 如果不存在则panic
func MustGetGlobalDBByDBName(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	db, ok := CLA_DBList[dbname]
	if !ok || db == nil {
		panic("db no init")
	}
	return db
}
