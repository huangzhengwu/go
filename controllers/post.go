package controllers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"web_app/logic"
	"web_app/models"
)

func CreatePostHandle(c *gin.Context) {
	// 1.获取参数以及数据校验
	p := new(models.Post)
	// c.ShouldBindJSON  //validator --> binging tag
	if err := c.ShouldBindJSON(&p); err != nil {
		zap.L().Error("c.ShouldBindJSON(&p)", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	// 从 c 获取当前登录用户的id
	userID, err := GetCurrentUser(c)
	if err != nil {
		zap.L().Error("GetCurrentUser(c)", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	p.AuthorID = userID
	// 2.创建帖子
	if err := logic.CreatePost(p); err != nil {
		zap.L().Error("logic.CreatePost(p)", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 3.返回响应
	ResponseSuccess(c, nil)
}
