package handler

import (
	"time"

	"github.com/gin-gonic/gin"

	"gin-framework/app/types"
	"gin-framework/app/utils"
	"gin-framework/config"
)

func LoginHandler(ctx *gin.Context) {
	var req types.UserReq
	if err := ctx.ShouldBind(&req); err != nil {
		utils.ValidatorError(ctx, err)
		return
	}
	token, err := utils.GenerateToken(222, config.GlobalConfig.Jwt.Secret)
	if err != nil {
		utils.Fail(ctx, "", nil)
		return
	}
	resp := types.UserResp{
		ID: 222,
		UserName: "yezi",
		CreateAt: time.Now().Unix(),
		Token: token,
	}
	utils.SuccessWithData(ctx, &resp)
}

func GetUserHandler(ctx *gin.Context) {
	// var req types.UserReq
	// if err := ctx.ShouldBind(&req); err != nil {
	// 	utils.ValidatorError(ctx, err)
	// } else {
		utils.Success(ctx)
	// }
}
