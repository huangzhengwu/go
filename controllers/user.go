package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"web_app/logic"
	"web_app/models"
)

func SignUpHandler(c *gin.Context) {
	//1.获取参数和参数校验
	p := new(models.ParamSignUp)
	if err := c.ShouldBindJSON(&p); err != nil {
		//参数有误
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors) //断言
		if !ok {
			ResponseError(c, CodeServerBusy)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	//fmt.Println(p)
	//2.业务处理
	if err := logic.SignUp(p); err != nil {
		ResponseError(c, CodeSignUpFailed)
		return
	}
	//3.返回响应
	ResponseSuccess(c, nil)
	return
}

func Login(c *gin.Context) {
	//1.获取参数和参数校验
	p := new(models.ParamLogin)
	if err := c.ShouldBindJSON(&p); err != nil {
		zap.L().Error("Login invalid param err:", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeServerBusy)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	//2.判断用户名密码是否错误
	token, err := logic.Login(p)
	if err != nil {
		zap.L().Error("Login invalid password err:", zap.Error(err))
		ResponseError(c, CodeInvalidPassword)
		return
	}
	//3.返回响应
	ResponseSuccess(c, token)
	return
}
