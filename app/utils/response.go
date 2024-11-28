package utils

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"gin-framework/pkg/e"
)

func ResponseJson(ctx *gin.Context, httpCode, dataCode int, msg string, data interface{}) {
	ctx.JSON(httpCode, gin.H{
		"code": dataCode,
		"msg":  msg,
		"data": data,
	})
}

func Success(ctx *gin.Context) {
	ResponseJson(ctx, e.SUCCESS, e.SUCCESS, e.CodeMsg[e.SUCCESS], map[string]interface{}{})
}

func SuccessWithMsg(ctx *gin.Context, msg string) {
	ResponseJson(ctx, e.SUCCESS, e.SUCCESS, msg, map[string]interface{}{})
}

func SuccessWithData(ctx *gin.Context, data interface{}) {
	ResponseJson(ctx, e.SUCCESS, e.SUCCESS, e.CodeMsg[e.SUCCESS], data)
}

func SuccessWithDetail(ctx *gin.Context, msg string, data interface{}) {
	ResponseJson(ctx, e.SUCCESS, e.SUCCESS, msg, data)
}

func Fail(ctx *gin.Context, dataCode int, msg string, data interface{}) {
	ResponseJson(ctx, e.ERROR, dataCode, msg, data)
}

func ValidatorError(ctx *gin.Context, err error) {
	if errs, ok := err.(validator.ValidationErrors); ok {
		// 翻译错误并移除结构体名称
		wrongParam := RemoveTopStruct(errs.Translate(Trans))
		// 返回结构体验证错误的响应
		ResponseJson(ctx, e.InvalidParams, e.InvalidParams, e.CodeMsg[e.InvalidParams], wrongParam)
	} else {
		errStr := err.Error()
		// multipart:nextpart:eof 错误表示验证器需要一些参数，但是调用者没有提交任何参数
		// 特定错误类型处理：multipart:nextpart:eof
		if strings.Contains("multipart:nextpart:eof", strings.ToLower(errStr)) {
			// 更明确的提示用户补充必填参数
			ResponseJson(ctx, e.InvalidParams, e.InvalidParams, e.CodeMsg[e.InvalidParams], gin.H{
				"tips": "该接口不允许所有参数都为空，请按照接口要求提交必填参数",
			})
		} else {
			// 对于其他错误，直接返回错误信息
			ResponseJson(ctx, e.InvalidParams, e.InvalidParams, e.CodeMsg[e.InvalidParams], gin.H{
				"tips": errStr,
			})
		}
	}
	ctx.Abort()
}
