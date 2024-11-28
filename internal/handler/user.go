package handler

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"gin-framework/pkg/e"
)

func GetUserHandler(ctx *gin.Context) {
	e.Success(ctx)
	zap.L().Warn("测试zap日志输出", zap.String("key1", "value1"))
}
