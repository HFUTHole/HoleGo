package controllers

import (
	"github.com/gin-gonic/gin"
	"hole/pkgs/common/response"
	"hole/pkgs/common/utils"
	"hole/pkgs/service"
	"strconv"
)

type ReplyReq struct {
	Text   string `json:"text"`
	Root   int64  `json:"root"`
	Parent int64  `json:"parent"`
	Real   bool   `json:"real"`
	Nick   string `json:"nick"`
	Avatar string `json:"avatar"`
}

func CreateReply() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cid, err := strconv.ParseInt(ctx.Param("cid"), 10, 64)
		if err != nil {
			response.Error403(ctx, "参数 :cid 解析错误")
			return
		}

		uid, err := utils.GetUid(ctx)
		if err != nil {
			response.Error403(ctx, "您可能还没有登录哦")
			return
		}

		var req ReplyReq
		err = ctx.BindJSON(&req)
		if err != nil {
			response.Error403(ctx, "body 解析错误")
			return
		}

		reply, err := service.CreateReply(uid, cid, req.Text, req.Nick, req.Avatar, req.Root, req.Parent, req.Real)

		if err != nil {
			response.HandleBusinessException(ctx, err)
			return
		}

		response.Success(ctx, reply)

	}
}

func GetReplyPage() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cid, err := strconv.ParseInt(ctx.Param("cid"), 10, 64)
		if err != nil {
			response.Error403(ctx, "参数 :cid 解析错误")
			return
		}
		pageSize, err := strconv.ParseInt(ctx.Query("pageSize"), 10, 64)
		if err != nil || pageSize >= 10 {
			pageSize = 10
		}

		maxId, err := strconv.ParseInt(ctx.Query("maxId"), 10, 64)
		if err != nil {
			maxId = 0
		}

		page, err := service.GetReplyPage(cid, int(pageSize), maxId)
		if err != nil {
			response.HandleBusinessException(ctx, err)
			return
		}

		response.Success(ctx, page)
	}
}
