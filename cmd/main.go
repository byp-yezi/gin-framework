package main

import (
	"gin-framework/initialize"
)

func main() {
	initialize.InitConfig()
	initialize.InitMysql()
	router := initialize.InitRouter()
	router.Run(initialize.GlobalConfig.Server.Port)
}
