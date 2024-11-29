package main

import (
	"go.uber.org/zap"
	
	"gin-framework/app/utils"
	"gin-framework/initialize"
	"gin-framework/config"
)

func main() {
	// 配置初始化
	config.InitConfig()

	// zap初始化
	initialize.InitZapLog()
	
	// mysql初始化
	initialize.InitMysql()

	// validator初始化
	if err := utils.InitTrans("zh"); err != nil {
		zap.S().Fatalf("init trans failed, err: ", err.Error())
	}

	// 路由初始化
	router := initialize.InitRouter()
	router.Run(config.GlobalConfig.Server.Port)
}
