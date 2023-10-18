package core

import (
	"Chinese_Learning_App/core/internal"
	"Chinese_Learning_App/global"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Viper
func Viper(path ...string) *viper.Viper {
	var config string

	// todo: 这里不做环境配置文件的选择

	config = internal.ConfigDefaultFile
	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("读取配置文件出错:%s \n", err))
	}

	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("配置文件更改:", e.Name)
		err = v.Unmarshal(&global.CLA_CONFIG)
		if err != nil {
			fmt.Println(err)
		}
	})
	if err = v.Unmarshal(&global.CLA_CONFIG); err != nil {
		fmt.Println(err)
	}
	//todo: root的适配性暂时不考虑
	return v
}
