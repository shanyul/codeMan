package app

import (
	"chujian-api/pkg/exception"
	"github.com/gin-gonic/gin"
)

type ResponseForm struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Response setting gin.JSON
func Response(ctx *gin.Context, httpCode, errCode int, data interface{}, msg string) {
	if msg == "" {
		msg = exception.GetMsg(errCode)
	}
	ctx.JSON(httpCode, ResponseForm{
		Code: errCode,
		Msg:  msg,
		Data: data,
	})
	return
}

// ResponseHtml setting gin.Html
func ResponseHtml(ctx *gin.Context, template string, httpCode, errCode int, data interface{}, msg string) {
	if msg == "" {
		msg = exception.GetMsg(errCode)
	}
	ctx.HTML(httpCode, template, ResponseForm{
		Code: errCode,
		Msg:  msg,
		Data: data,
	})
	return
}
