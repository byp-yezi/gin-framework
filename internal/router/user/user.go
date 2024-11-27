package user

import (
	"github.com/gin-gonic/gin"

	"gin-framework/internal/handler"
)


func InitUserGroup(r *gin.RouterGroup) (router gin.IRoutes) {
	userGroup := r.Group("user")
	{
		userGroup.GET("", handler.GetUserHandler)
	}
	return userGroup
}
