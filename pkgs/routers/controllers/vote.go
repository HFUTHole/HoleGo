package controllers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"hole/pkgs/common"
	"hole/pkgs/common/response"
	"hole/pkgs/config/logger"
	"hole/pkgs/service"
	"time"
)

type CreateVoteReq struct {
	Uid     int64           `form:"uid" binding:"required"`
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

		voting, err := service.CreateContentVoting(vote.Uid, vote.Cid, vote.Options, time.Time(vote.EndTime))
		if err != nil {
			response.HandleBusinessException(ctx, err)
			return
		}

		response.Success(ctx, voting)
	}
}

type VoteReq struct {
	Uid int64 `json:"uid" binding:"required"`
	Cid int64 `json:"cid" binding:"required"`
	Vid int64 `json:"vid" binding:"required"`
}

func Vote() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var vote VoteReq
		err := ctx.BindJSON(&vote)
		if err != nil {
			response.Error403(ctx, "参数错误")
			return
		}

		createVote, err := service.CreateVote(vote.Uid, vote.Cid, vote.Vid)
		if err != nil {
			response.HandleBusinessException(ctx, err)
			return
		}

		response.Success(ctx, createVote)

	}
}
