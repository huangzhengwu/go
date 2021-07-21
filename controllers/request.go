package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
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
