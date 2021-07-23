package controllers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
	"web_app/logic"
	"web_app/models"
)

// CreatePostHandle 创建帖子
func CreatePostHandle(c *gin.Context) {
	// 1.获取参数以及数据校验
	p := new(models.Post)
	// c.ShouldBindJSON  //validator --> binging tag
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Debug("c.ShouldBindJSON(&p)", zap.Any("err", err))
		zap.L().Error("create post with invalid param", zap.Error(err))
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

// PostDetailHandle 帖子详情
func PostDetailHandle(c *gin.Context) {
	// 1.获取参数
	pIdStr := c.Param("id")
	pId, err := strconv.ParseInt(pIdStr, 10, 64)
	if err != nil {
		zap.L().Error("CodeInvalidParam err", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	// 2.获取帖子详情
	data, err := logic.GetPostDetail(pId)
	if err != nil {
		zap.L().Error("PostDetailHandle err", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 3.返回数据
	ResponseSuccess(c, data)
}

// PostListHandle 帖子列表
func PostListHandle(c *gin.Context) {
	// 1.接收参数并校验
	pageStr := c.Query("page")
	sizeStr := c.Query("size")
	page, err := strconv.ParseInt(pageStr, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	size, err := strconv.ParseInt(sizeStr, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	// 2.获取帖子数据列表
	data, err := logic.GetPostList(page, size)
	if err != nil {
		zap.L().Error("controllers logic.GetPostList() err", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 3.返回数据
	ResponseSuccess(c, data)
}

// PostListRedisHandle 帖子列表(使用redis)
// 按照创建时间排序 或者 按照 分数排序
// 1.获取参数
// 2.去redis查询id列表
// 3.根据id去数据库查询帖子详情信息
func PostListRedisHandle(c *gin.Context) {
	// GET请求参数(query string)： /api/v1/post-list?page=1&size=2&order=time
	// 1.接收参数并校验
	// 初始化结构体时指定初始参数
	p := &models.ParamPostList{
		Page:  1,
		Size:  10,
		Order: models.OrderTime, // magic string
	}
	//c.ShouldBind() //根据请求中的数据类型选择相应的方法获取数据
	//c.ShouldBindJSON() //如果请求中携带的时json数据，才能用这个方法获取数据
	if err := c.ShouldBindQuery(p); err != nil {
		zap.L().Error("c.ShouldBindQuery(p) err： ", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	// 2.获取帖子数据列表
	data, err := logic.GetPostListRedis(p)
	if err != nil {
		zap.L().Error("controllers PostListRedisHandle logic.GetPostList() err", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 3.返回数据
	ResponseSuccess(c, data)
}

func GetCommunityPostListHandle(c *gin.Context) {
	p := &models.ParamCommunityPostList{
		ParamPostList: &models.ParamPostList{
			Page:  1,
			Size:  10,
			Order: models.OrderTime,
		},
	}
	//c.ShouldBind() //根据请求中的数据类型选择相应的方法获取数据
	//c.ShouldBindJSON() //如果请求中携带的时json数据，才能用这个方法获取数据
	if err := c.ShouldBindQuery(p); err != nil {
		zap.L().Error("GetCommunityPostListHandle c.ShouldBindQuery(p) err： ", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	// 2.获取帖子数据列表
	data, err := logic.GetCommunityPostListRedis(p)
	if err != nil {
		zap.L().Error("controllers logic.GetPostList() err", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 3.返回数据
	ResponseSuccess(c, data)
}
