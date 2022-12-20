package controllers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"hole/pkgs/common/response"
	"hole/pkgs/config/fileservice"
	"hole/pkgs/config/logger"
	"strings"
)

func UpdateImage() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		file, err := ctx.FormFile("image")
		filename := file.Filename

		suffix := filename[strings.LastIndex(filename, ".")+1:]

		logger.GetLogger().Info("update filename", zap.String("suffix", suffix))

		if err != nil {
			logger.GetLogger().Fatal(err.Error(), zap.Error(err))
		}

		src, err := file.Open()
		id, err := fileservice.PutFile(src, file.Size, "image/png")

		if err != nil {
			logger.GetLogger().Error("error", zap.Error(err))
		}

		ctx.JSON(200, gin.H{
			"url": "http://127.0.0.1:8888/image/temp/" + id,
			"id":  id,
		})
	}
}

func DownloadImage() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		bucket := ctx.Param("bucket")
		id := ctx.Param("id")
		logger.GetLogger().Info("path param", zap.String("bucket", bucket), zap.String("id", id))

		if bucket == fileservice.TempBucket {
			obj, err := fileservice.GetTemp(id)
			if err != nil {
				response.Error500(ctx, err.Error())
				return
			}
			response.WriteObject(ctx, obj)
			return
		}

		if bucket == fileservice.ContentBucket {
			obj, err := fileservice.GetContent(id)
			if err != nil {
				response.Error500(ctx, err.Error())
				return
			}
			response.WriteObject(ctx, obj)
			return
		}

		if bucket == fileservice.AvatarBucket {
			obj, err := fileservice.GetAvatar(id)
			if err != nil {
				response.Error500(ctx, err.Error())
				return
			}
			response.WriteObject(ctx, obj)
			return
		}

		response.Error404(ctx, "访问资源不存在")
	}
}