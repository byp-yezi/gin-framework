package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"

	"gin-framework/app/utils"
	"gin-framework/config"
)

func Jwt() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader(config.GlobalConfig.Jwt.TokenKey)
		if token == "" {
			utils.UnauthorizedException(ctx, "")
			ctx.Abort()
			return
		}
		token = strings.TrimSpace(token)
		if _, err := utils.ParseToken(token, config.GlobalConfig.Jwt.Secret); err != nil {
			utils.UnauthorizedException(ctx, err.Error())
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
