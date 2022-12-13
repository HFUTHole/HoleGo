package response

import (
	"hole/src/exception"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Data struct {
	Code    int         `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data,omitempty"` // omitempty当data为空时,不展示这个字段
}

func Error403(ctx *gin.Context, msg string) {
	Error(ctx, http.StatusForbidden, msg)
}

func Error500(ctx *gin.Context, msg string) {
	Error(ctx, http.StatusInternalServerError, msg)
}

func Error(ctx *gin.Context, code int, msg string) {
	rd := &Data{
		Code:    code,
		Message: msg,
		Data:    nil,
	}
	ctx.JSON(code, rd)
}

func Success(ctx *gin.Context, data interface{}) {
	rd := &Data{
		Code:    http.StatusOK,
		Message: http.StatusText(http.StatusOK),
		Data:    data,
	}
	ctx.JSON(http.StatusOK, rd)
}

func SuccessWithMsg(ctx *gin.Context, data interface{}, msg string) {
	rd := &Data{
		Code:    http.StatusOK,
		Message: msg,
		Data:    data,
	}
	ctx.JSON(http.StatusOK, rd)
}

func HandleBusinessException(ctx *gin.Context, err error) {
	var e interface{}
	e = &err
	client, ok := e.(*exception.ClientException)
	if ok {
		Error403(ctx, client.Error())
		return
	}

	server, ok := e.(*exception.ServerException)
	if ok {
		Error500(ctx, server.Error())
		return
	}

	business, ok := e.(*exception.BusinessException)
	if ok {
		Error500(ctx, business.Error())
		return
	}

	Error500(ctx, "未知异常")
}
