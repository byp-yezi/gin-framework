package initialize

import (
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"gin-framework/internal/router/user"
)

func InitRouter() *gin.Engine {
	if GlobalConfig.Server.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()
	r.Use(ginzap.Ginzap(zap.L(), time.RFC3339, false))
	r.Use(ginzap.RecoveryWithZap(zap.L(), false))
	// 路由分组
	var (
		userGroup = r.Group("/")
	)
	user.InitUserGroup(userGroup)
	return r
}
