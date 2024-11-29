package initialize

import (
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"gin-framework/app/router/user"
	"gin-framework/config"
	"gin-framework/middleware"
)

func InitRouter() *gin.Engine {
	if config.GlobalConfig.Server.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(ginzap.Ginzap(zap.L(), time.RFC3339, false))
	r.Use(ginzap.RecoveryWithZap(zap.L(), false))
	r.Use(middleware.Cors())
	// 路由分组
	var (
		userGroup = r.Group("/user")
	)
	user.InitUserGroup(userGroup)
	return r
}
