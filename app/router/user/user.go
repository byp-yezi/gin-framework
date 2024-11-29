package user

import (
	"github.com/gin-gonic/gin"

	"gin-framework/app/handler"
	"gin-framework/middleware"
)


func InitUserGroup(r *gin.RouterGroup) (router gin.IRoutes) {
	r.POST("login", handler.LoginHandler)
	r.Use(middleware.Jwt())
	{
		r.GET("get", handler.GetUserHandler)
	}
	return r
}
