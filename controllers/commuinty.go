package controllers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
	"web_app/logic"
)

//   -------  跟社区相关   --------

// CommunityHandle 获取社区列表
func CommunityHandle(c *gin.Context) {
	//查询所有社区 （community_id ,community_name）  以列表形式返回
	date, err := logic.GetCommunityList()

	if err != nil {
		zap.L().Error("login.GetCommunityHandle() err ", zap.Error(err))
		ResponseError(c, CodeServerBusy) //不轻易把服务端报错暴露给外面
		return
	}
	ResponseSuccess(c, date)
}

// CommunityDetailHandle 根据社区id，获取详情数据
func CommunityDetailHandle(c *gin.Context) {
	// 获取参与id 并验证id有效性
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	//根据id查询社区详情数据
	data, err := logic.GetCommunityDetail(id)
	if err != nil {
		zap.L().Error("login.GetCommunityDetail() err ", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}
