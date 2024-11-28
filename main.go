package main

import (
	"fmt"
	"gin-framework/initialize"

	"go.uber.org/zap"
)

func main() {
	initialize.InitConfig()

	switch initialize.GlobalConfig.Server.Mode {
	case "debug":
		initialize.InitZapLog(zap.DebugLevel, initialize.GlobalConfig.Log.LogformatConsole)
		fmt.Println("1111111111111")
	case "release":
		initialize.InitZapLog(zap.InfoLevel, initialize.GlobalConfig.Log.LogformatConsole)
		fmt.Println("2222222222222")
	default:
		initialize.InitZapLog(zap.InfoLevel, initialize.GlobalConfig.Log.LogformatConsole)
		fmt.Println("3333333333333")
	}
	
	initialize.InitMysql()
	router := initialize.InitRouter()
	router.Run(initialize.GlobalConfig.Server.Port)
}
