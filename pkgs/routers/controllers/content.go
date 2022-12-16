package controllers

import (
	"github.com/gin-gonic/gin"
	"hole/pkgs/common/response"
	"hole/pkgs/service"
	"strconv"
)

type ContentReq struct {
	Uid       int64    `json:"uid"`
	Title     string   `json:"title" binding:"required"`
	Message   string   `json:"message" binding:"required"`
	Tags      []string `json:"tags"`
	ImageUrls []string `json:"images"`
	Real      bool     `json:"real"`
}

func CreateContent() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req ContentReq
		err := ctx.BindJSON(&req)
		if err != nil {
			response.Error403(ctx, "参数错误")
			return
		}
		if len(req.Title) <= 4 {
			response.Error403(ctx, "帖子标题太少了")
			return
		}
		if len(req.Title) >= 32 {
			response.Error403(ctx, "帖子标题超过最大长度")
			return
		}

		if len(req.Message) <= 10 {
			response.Error403(ctx, "帖子内容太少了")
			return
		}

		if len(req.Message) >= 1024 {
			response.Error403(ctx, "帖子内容超过最大长度")
			return
		}
		if len(req.Tags) >= 5 {
			response.Error403(ctx, "帖子最多有五个标签哦")
			return
		}
		if len(req.ImageUrls) >= 3 {
			response.Error403(ctx, "帖子最多有三张照片哦")
			return
		}

		// TODO( Uid ctx)
		content, e := service.CreateContent(req.Uid, req.Title, req.Message, req.Tags, req.ImageUrls, req.Real)
		if e != nil {
			response.HandleBusinessException(ctx, e)
			return
		}

		response.Success(ctx, content)
	}
}

type PageInfo struct {
	MaxId    int64 `json:"maxId" form:"maxId"`
	PageSize int   `json:"pageSize" form:"pageSize"`
}

func GetContentPage() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var page PageInfo
		_ = ctx.BindQuery(&page)

		if page.PageSize == 0 {
			page.PageSize = 20
		}

		if page.MaxId > 0 {
			nextPage, err := service.GetContentNextPage(page.MaxId, page.PageSize)
			if err != nil {
				response.HandleBusinessException(ctx, err)
				return
			}
			response.Success(ctx, nextPage)
			return
		}

		contentPage, err := service.GetContentPage(page.PageSize)
		if err != nil {
			response.HandleBusinessException(ctx, err)
			return
		}
		response.Success(ctx, contentPage)
	}
}

func GetContent() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cid, err := strconv.ParseInt(ctx.Query("cid"), 10, 64)
		if err != nil {
			response.Error403(ctx, "参数错误: cid")
		}

		content, err := service.GetContent(cid)
		if err != nil {
			response.HandleBusinessException(ctx, err)
			return
		}

		response.Success(ctx, content)
	}
}
