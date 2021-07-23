package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
	"web_app/middlewares"
)

var ErrorUserNotLogin = errors.New("用户没有登录")

func GetCurrentUser(c *gin.Context) (userId int64, err error) {
	uid, ok := c.Get(middlewares.CtxUserIDKey)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	userId, ok = uid.(int64)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	return
}

func getPageInfo(c *gin.Context) (int64, int64) {
	pageStr := c.Query("page")
	sizeStr := c.Query("size")
	var (
		page int64
		size int64
		err  error
	)

	page, err = strconv.ParseInt(pageStr, 10, 64)
	if err != nil {
		page = 1
	}
	size, err = strconv.ParseInt(sizeStr, 10, 64)
	if err != nil {
		size = 2
	}
	return page, size
}
