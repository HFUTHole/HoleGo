package service

import (
	"github.com/bwmarrin/snowflake"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"hole/pkgs/common/utils"
	"hole/pkgs/config/fileservice"
	"hole/pkgs/config/logger"
	"hole/pkgs/config/mysql"
	"hole/pkgs/dao"
	"hole/pkgs/exception"
	"hole/pkgs/models"
	"hole/pkgs/models/role"
	"hole/pkgs/models/vo"
)

var ID, _ = snowflake.NewNode(1)

func CreateReply(uid int64, cid int64, text string, nick string, avatar string, root int64, parent int64, real bool) ([]vo.ReplyVO, error) {

	scapeText := utils.Scape(text)
	if len(scapeText) > 1024-128 {
		return nil, &exception.ClientException{Msg: "评论内容过长"}
	}

	err := mysql.GetDB().Transaction(func(tx *gorm.DB) error {
		user, err := dao.GetUserByID(tx, uid)
		if err != nil {
			return &exception.ClientException{Msg: "用户不存在"}
		}
		if user.Role.Validate(role.NormalUserRole, role.AdminRole, role.SuperUserRole) {
			return &exception.ClientException{Msg: "您还不可以回复哦"}
		}

		content, err := dao.GetContent(tx, cid)
		if err != nil {
			return &exception.ClientException{Msg: "未查询到帖子"}
		}

		if root != -1 {
			_, e := dao.GetRootReply(tx, cid, root)
			if e != nil {
				return &exception.ClientException{Msg: "没有查询到根回复"}
			}
		}
		var parentReply *models.Reply

		if parent != -1 {
			var e error
			parentReply, e = dao.GetParentReply(tx, cid, parent)
			if e != nil {
				return &exception.ClientException{Msg: "没有查询到父回复"}
			}
		}

		var aid int64

		if !real {
			if len(nick) > 32 {
				return &exception.ClientException{Msg: "nick 过长"}
			}
			aliasID, e := utils.AliasID(nick)
			if e != nil {
				return &exception.ClientException{Msg: e.Error()}
			}
			aid = aliasID

			if !fileservice.Exists(fileservice.TempBucket, avatar) {
				return &exception.ClientException{Msg: "头像不存在"}
			}
		} else {
			nick = user.Username
			avatar = user.Avatar
			aid = user.ID
		}

		rid := ID.Generate().Int64()

		var atName int
		if root != -1 && root != parent {

			at := &models.AtName{
				ReplyID: rid,
				Uid:     parentReply.Aid,
				Text:    "@" + parentReply.Nick,
			}
			e := dao.CreateAtName(tx, at)
			if e != nil {
				logger.GetLogger().Error("create @nick failure",
					zap.Int64("uid", uid),
					zap.Any("data", at),
					zap.Error(e),
				)
				//return &exception.ServerException{Msg: "回复失败"}
			}

			scapeText = utils.Scape("回复 @" + parentReply.Nick + " " + scapeText)
			atName = 1
		}

		reply := &models.Reply{
			ID:      rid,
			Cid:     content.ID,
			Root:    root,
			Parent:  parent,
			Uid:     user.ID,
			Aid:     aid,
			Real:    utils.BoolToInt(real),
			Nick:    nick,
			Avatar:  avatar,
			Message: scapeText,
			AtName:  atName,
		}

		err = dao.CreateReply(tx, reply)
		if err != nil {
			logger.GetLogger().Error("create @nick failure",
				zap.Int64("uid", uid),
				zap.Any("data", reply))
			return &exception.ClientException{Msg: "回复失败 "}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return GetReplyPage(cid, 10, 0)
}

func GetReplyPage(cid int64, pageSize int, maxId int64) ([]vo.ReplyVO, error) {

	var replyVo []vo.ReplyVO
	err := mysql.GetDB().Transaction(func(tx *gorm.DB) error {
		var replies []models.Reply
		if maxId <= 0 {
			page, err := dao.GetRootReplies(tx, cid, pageSize)
			if err != nil {
				return &exception.ClientException{Msg: "未查询到回复"}
			}
			replies = page
		} else {
			next, err := dao.GetRootReplyNext(tx, cid, pageSize, maxId)
			if err != nil {
				return err
			}
			replies = next
		}

		replyVo = make([]vo.ReplyVO, len(replies))
		for i, reply := range replies {

			replyVo[i] = vo.ReplyVO{
				ID:        reply.ID,
				Cid:       reply.Cid,
				Uid:       reply.Aid,
				Nick:      reply.Nick,
				Avatar:    reply.Avatar,
				Message:   reply.Message,
				CreatedAt: reply.CreatedAt,
			}

			if reply.AtName != 0 {
				atNames, _ := dao.GetAtNameByRid(tx, reply.ID)
				atNameVOS := make(map[string]vo.AtNameVO, len(atNames))

				for _, atName := range atNames {
					atNameVOS[atName.Text] = vo.AtNameVO{
						ID:      atName.ID,
						ReplyID: atName.ReplyID,
						Text:    atName.Text,
						Uid:     atName.Uid,
					}
				}
				replyVo[i].AtName = atNameVOS
			}

			children, _ := dao.GetChildren(tx, cid, reply.ID)

			replyVo[i].List = make([]vo.ReplyChildVO, len(children))

			for i2, child := range children {
				replyVo[i].List[i2] = vo.ReplyChildVO{
					ID:        child.ID,
					Cid:       child.Cid,
					Root:      child.Root,
					Parent:    child.Parent,
					Uid:       child.Aid,
					Nick:      child.Nick,
					Avatar:    child.Avatar,
					Message:   child.Message,
					CreatedAt: child.CreatedAt,
				}

				if child.AtName != 0 {
					atNames, _ := dao.GetAtNameByRid(tx, child.ID)
					atNameVOS := make(map[string]vo.AtNameVO, len(atNames))

					for _, atName := range atNames {
						atNameVOS[atName.Text] = vo.AtNameVO{
							ID:      atName.ID,
							ReplyID: atName.ReplyID,
							Text:    atName.Text,
							Uid:     atName.Uid,
						}
					}
					replyVo[i].List[i2].AtName = atNameVOS
				}
			}
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return replyVo, nil
}
