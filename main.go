package main

import (
	"gin-framework/app/utils"
	"gin-framework/initialize"

	"go.uber.org/zap"
)

func main() {
	// 配置初始化
	initialize.InitConfig()

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
	router.Run(initialize.GlobalConfig.Server.Port)
}
