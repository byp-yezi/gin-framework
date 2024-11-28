package user

import (
	"github.com/gin-gonic/gin"

	"gin-framework/app/handler"
)


func InitUserGroup(r *gin.RouterGroup) (router gin.IRoutes) {
	userGroup := r.Group("user")
	{
		userGroup.POST("", handler.GetUserHandler)
	}
	return userGroup
}
