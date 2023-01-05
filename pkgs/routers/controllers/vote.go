package controllers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"hole/pkgs/common"
	"hole/pkgs/common/response"
	"hole/pkgs/common/utils"
	"hole/pkgs/config/logger"
	"hole/pkgs/service"
	"strconv"
	"time"
)

type CreateVoteReq struct {
	Cid     int64           `form:"cid" binding:"required"`
	Options []string        `form:"options" binding:"required"`
	EndTime common.DateTime `form:"endTime" time_format:"yyyy-MM-dd HH:mm:ss"`
}

func CreateContentVote() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var vote CreateVoteReq
		err := ctx.ShouldBind(&vote)
		if err != nil {
			logger.GetLogger().Error("da", zap.Error(err))
			response.Error403(ctx, "参数错误")
			return
		}

		uid, err := utils.GetUid(ctx)
		if err != nil {
			response.Error403(ctx, "你可能还没有登录")
			return
		}

		voting, err := service.CreateContentVoting(uid, vote.Cid, vote.Options, time.Time(vote.EndTime))
		if err != nil {
			response.HandleBusinessException(ctx, err)
			return
		}

		response.Success(ctx, voting)
	}
}

type DeleteContentVoteReq struct {
	Cid int64 `json:"cid" binding:"required"`
}

func DeleteContentVote() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		var req DeleteContentVoteReq
		err := ctx.BindJSON(&req)
		if err != nil {
			response.Error403(ctx, "参数错误")
			return
		}
		uid, err := utils.GetUid(ctx)
		if err != nil {
			response.HandleBusinessException(ctx, err)
			return
		}

		vote, err := service.DeleteContentVote(uid, req.Cid)
		if err != nil {
			response.HandleBusinessException(ctx, err)
			return
		}

		response.Success(ctx, vote)
	}
}

type VoteReq struct {
	Vid int64 `json:"vid" binding:"required"`
}

func CreateVote() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var vote VoteReq
		err := ctx.BindJSON(&vote)
		if err != nil {
			response.Error403(ctx, "参数错误")
			return
		}
		cid, err := strconv.ParseInt(ctx.Param("cid"), 10, 64)
		if err != nil {
			response.Error403(ctx, "参数错误 :cid")
		}

		uid, err := utils.GetUid(ctx)
		if err != nil {
			response.Error401(ctx, "您还没有登录哦")
		}

		createVote, err := service.CreateVote(uid, cid, vote.Vid)
		if err != nil {
			response.HandleBusinessException(ctx, err)
			return
		}

		response.Success(ctx, createVote)

	}
}

func CancelVote() gin.HandlerFunc {
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

		vo, err := service.CancelVote(uid, cid)

		if err != nil {
			response.HandleBusinessException(ctx, err)
			return
		}

		response.Success(ctx, vo)
	}
}
