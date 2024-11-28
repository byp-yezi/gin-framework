package e

import "github.com/gin-gonic/gin"

func ResponseJson(ctx *gin.Context, httpCode, dataCode int, msg string, data interface{}) {
	ctx.JSON(httpCode, gin.H{
		"code": dataCode,
		"msg": msg,
		"data": data,
	})
}

func Success(ctx *gin.Context) {
	ResponseJson(ctx, SUCCESS, SUCCESS, CodeMsg[SUCCESS], map[string]interface{}{})
}

func SuccessWithMsg(ctx *gin.Context, msg string) {
	ResponseJson(ctx, SUCCESS, SUCCESS, msg, map[string]interface{}{})
}

func SuccessWithData(ctx *gin.Context, data interface{}) {
	ResponseJson(ctx, SUCCESS, SUCCESS, CodeMsg[SUCCESS], data)
}

func SuccessWithDetail(ctx *gin.Context, msg string, data interface{}) {
	ResponseJson(ctx, SUCCESS, SUCCESS, msg, data)
}

func Fail(ctx *gin.Context, dataCode int, msg string, data interface{}) {
	ResponseJson(ctx, ERROR, dataCode, msg, data)
}
