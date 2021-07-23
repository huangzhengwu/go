package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResponseData struct {
	Code ResCode     `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func ResponseError(c *gin.Context, code ResCode) {
	res := &ResponseData{
		Code: code,
		Msg:  code.Msg(),
		Data: nil,
	}
	c.JSON(http.StatusOK, res)
}

func ResponseErrorWithMsg(c *gin.Context, code ResCode, msgStr interface{}) {
	res := &ResponseData{
		Code: code,
		Msg:  msgStr,
		Data: nil,
	}
	c.JSON(http.StatusOK, res)
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	res := &ResponseData{
		Code: CodeSuccess,
		Msg:  CodeSuccess.Msg(),
		Data: data,
	}
	c.JSON(http.StatusOK, res)
}
