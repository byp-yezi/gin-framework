package initialize

import (
	"gin-framework/internal/router/user"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	
	// 路由分组
	var (
		userGroup = r.Group("/")
	)
	user.InitUserGroup(userGroup)
	return r
}
