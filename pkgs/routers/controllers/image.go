package controllers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"hole/pkgs/common/response"
	"hole/pkgs/common/utils"
	"hole/pkgs/config/base"
	"hole/pkgs/config/fileservice"
	"hole/pkgs/config/logger"
)

func UpdateImage() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		file, err := ctx.FormFile("image")
		if err != nil {
			logger.GetLogger().Error("error form", zap.Error(err))
			response.Error403(ctx, "错误的表单信息")
			return
		}
		filename := file.Filename
		contentType, err := utils.FileNameToContentType(filename)

		if err != nil {
			response.Error403(ctx, err.Error())
			return
		}

		src, err := file.Open()
		id, err := fileservice.PutFile(src, file.Size, contentType)

		if err != nil {
			logger.GetLogger().Error("error", zap.Error(err))
		}

		ctx.JSON(200, gin.H{
			"url": base.GetDomain() + "/image/temp/" + id,
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
