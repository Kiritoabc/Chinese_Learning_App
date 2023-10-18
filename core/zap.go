package core

import (
	"Chinese_Learning_App/core/internal"
	"Chinese_Learning_App/global"
	"Chinese_Learning_App/utils"
	"fmt"
	"os"

	"go.uber.org/zap/zapcore"

	"go.uber.org/zap"
)

func Zap() (logger *zap.Logger) {
	if ok, _ := utils.PathExists(global.CLA_CONFIG.Zap.Director); !ok {
		fmt.Printf("create %v directory \n", global.CLA_CONFIG.Zap.Director)
		err := os.Mkdir(global.CLA_CONFIG.Zap.Director, os.ModePerm)
		if err != nil {
			return nil
		}
	}
	cores := internal.Zap.GetZapCores()
	logger = zap.New(zapcore.NewTee(cores...))

	if global.CLA_CONFIG.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}
