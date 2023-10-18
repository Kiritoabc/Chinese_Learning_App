package global

import (
	"Chinese_Learning_App/config"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// 定义一些全局变量
var (
	CLA_VP     *viper.Viper  //
	CLA_CONFIG config.Config // 全局的配置文件
	CLA_LOG    *zap.Logger
)
