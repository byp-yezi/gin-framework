package handler

import (
	"github.com/gin-gonic/gin"

	"gin-framework/app/types"
	"gin-framework/app/utils"
)

func GetUserHandler(ctx *gin.Context) {
	var req types.UserReq
	if err := ctx.ShouldBind(&req); err != nil {
		utils.ValidatorError(ctx, err)
	} else {
		utils.Success(ctx)
	}
}
