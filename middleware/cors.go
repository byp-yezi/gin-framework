package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	defaultConfig := cors.DefaultConfig()
	defaultConfig.AllowAllOrigins = true
	defaultConfig.AllowCredentials = true
	// defaultConfig.AddAllowHeaders("DNT", "X-Mx-ReqToken", "X-CSRF-Token", "Keep-Alive", "User-Agent", "X-Requested-With", "X-Request-ID", "If-Modified-Since", "Cache-Control", "Authorization")
	return cors.New(defaultConfig)
}
