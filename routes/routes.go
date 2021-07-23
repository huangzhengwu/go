package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web_app/controllers"
	"web_app/logger"
	"web_app/middlewares"
)

func Setup(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) //设置成发布模式
	}
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	v1 := r.Group("/api/v1")

	r.POST("/signup", controllers.SignUpHandler)
	r.POST("/login", controllers.Login)
	v1.Use(middlewares.JWTAuthMiddleware()) //应用 JWT认证 中间件
	{
		//帖子社区
		v1.GET("/community", controllers.CommunityHandle)
		v1.GET("/community/:id", controllers.CommunityDetailHandle)
		//帖子curd
		v1.POST("/add-post", controllers.CreatePostHandle)
		v1.GET("/post-detail/:id", controllers.PostDetailHandle)
		v1.GET("/post-list", controllers.PostListHandle)
		v1.GET("/post-list—redis", controllers.PostListRedisHandle)

		//投票
		v1.POST("/vote", controllers.PostVoteController)

	}
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
