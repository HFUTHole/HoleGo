package response

import (
	"github.com/minio/minio-go/v6"
	"go.uber.org/zap"
	"hole/pkgs/config/logger"
	"hole/pkgs/exception"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Data struct {
	Code    int         `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data,omitempty"` // omitempty当data为空时,不展示这个字段
}

func Error401(ctx *gin.Context, msg string) {
	Error(ctx, http.StatusUnauthorized, msg)
}

func Error403(ctx *gin.Context, msg string) {
	Error(ctx, http.StatusForbidden, msg)
}

func Error404(ctx *gin.Context, msg string) {
	Error(ctx, http.StatusNotFound, msg)
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

func WriteObject(ctx *gin.Context, obj *minio.Object) {

	stat, err := obj.Stat()
	if err != nil {
		Error404(ctx, "资源不存在")
		return
	}
	contentType := stat.ContentType
	ctx.Header("ContentType", contentType)
	writer := ctx.Writer

	var buf = make([]byte, 1024)
	n, err := obj.Read(buf)
	for ; err == nil; n, err = obj.Read(buf) {
		if n == 1024 {
			_, e := writer.Write(buf)
			if e != nil {
				logger.GetLogger().Error("图片写回错误", zap.Error(e))
			}
		} else {
			_, e := writer.Write(buf[0:n])
			if e != nil {
				logger.GetLogger().Error("图片写回错误", zap.Error(e))
			}
		}
	}
	writer.Flush()
}

func HandleBusinessException(ctx *gin.Context, err error) {

	client, ok := err.(*exception.ClientException)
	if ok {
		Error403(ctx, client.Error())
		return
	}

	server, ok := err.(*exception.ServerException)
	if ok {
		Error500(ctx, server.Error())
		return
	}

	business, ok := err.(*exception.BusinessException)
	if ok {
		Error500(ctx, business.Error())
		return
	}

	Error500(ctx, "未知异常")
}
