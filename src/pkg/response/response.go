package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/**
 * @Author huchao
 * @Description
 * @Date 22:11 2022/2/10
 **/

type Data struct {
	Code    MyCode      `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data,omitempty"` // omitempty当data为空时,不展示这个字段
}

func Error(ctx *gin.Context, msg string) {
	rd := &Data{
		Code:    http.StatusInternalServerError,
		Message: msg,
		Data:    nil,
	}
	ctx.JSON(http.StatusInternalServerError, rd)
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
