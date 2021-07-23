package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"web_app/logic"
	"web_app/models"
)

//投票

func PostVoteController(c *gin.Context) {
	//参数校验
	p := new(models.ParamVoteData)
	if err := c.ShouldBindJSON(p); err != nil {
		errs, ok := err.(validator.ValidationErrors) //类型断言
		if !ok {
			zap.L().Error(" c.ShouldBindJSON(p) err", zap.Error(err))
			ResponseError(c, CodeInvalidParam)
			return
		}
		errData := removeTopStruct(errs.Translate(trans)) //去除提示信息中的结构体名称
		zap.L().Error(" c.ShouldBindJSON(p) err", zap.Error(err))
		ResponseErrorWithMsg(c, CodeInvalidParam, errData)
		return
	}
	//获取当前用户id
	userID, err := GetCurrentUser(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
	}
	if err := logic.VoteForPost(userID, p); err != nil {
		zap.L().Error("logic.VoteForPost(userID,p) err ", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, nil)
}
