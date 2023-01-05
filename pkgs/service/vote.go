package service

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"hole/pkgs/common/utils"
	"hole/pkgs/config/logger"
	"hole/pkgs/config/mysql"
	"hole/pkgs/dao"
	"hole/pkgs/exception"
	"hole/pkgs/models/role"
	"hole/pkgs/models/vo"
	"time"
)

func CreateContentVoting(uid int64, cid int64, options []string, endTime time.Time) (*vo.ContentVO, error) {
	options = utils.ScapeSlice(options)
	if !utils.SliceElementMaxLength(options, 64) {
		return nil, &exception.ClientException{Msg: "创建投票选项超过了最大长度"}
	}

	if endTime.Before(time.Now()) {
		return nil, &exception.ClientException{Msg: "投票时间必须大于当前时间"}
	}

	err := mysql.GetDB().Transaction(func(tx *gorm.DB) error {

		content, e := GetContent(cid)
		if e != nil && content.ID == 0 {
			return &exception.ClientException{Msg: "未查询到该贴子信息"}
		}

		user, e := dao.GetUserByID(tx, uid)
		if e != nil && user.ID == 0 {
			return &exception.ClientException{Msg: "没有查询到用户信息"}
		}

		if !(user.Role.Validate(role.NormalUserRole) && user.ID != content.Uid) {
			return &exception.ClientException{Msg: "您不可以发起投票哦"}
		}

		count, e := dao.GetCountContentByCid(tx, cid)
		if e != nil || count > 0 {
			return &exception.ClientException{Msg: "已创建投票"}
		}

		e = dao.SetContentVoteEndTime(tx, cid, 1, endTime)
		if e != nil {
			return &exception.ServerException{Msg: "创建投票错误"}
		}

		e = dao.CreateVote(tx, cid, options)
		if e != nil {
			return &exception.ServerException{Msg: "创建投票错误"}
		}

		return nil
	})

	if err != nil {
		logger.GetLogger().Error(err.Error(), zap.Error(err), zap.Int64("uid", uid))
		return nil, err
	}

	return GetContent(cid)
}

func GetContentVoting(cid int64) ([]vo.VotingOptionVO, error) {
	db := mysql.GetDB()
	voting, err := dao.GetContentVoting(db, cid)

	if err != nil {
		return nil, &exception.ClientException{Msg: "帖子投票查询错误"}
	}

	return vo.ConvertVotingOption(voting), nil
}

func CreateVote(uid int64, cid int64, vid int64) (*vo.ContentVO, error) {

	err := mysql.GetDB().Transaction(func(tx *gorm.DB) error {
		content, e := dao.GetContent(tx, cid)
		if e != nil && content.ID == 0 {
			return &exception.ClientException{Msg: "未查询到该贴子信息"}
		}

		if content.EndTime.Before(time.Now()) {
			return &exception.ClientException{Msg: "投票已结束"}
		}

		user, e := dao.GetUserByID(tx, uid)
		if e != nil && user.ID == 0 {
			return &exception.ClientException{Msg: "没有查询到用户信息"}
		}

		if !user.Role.Validate(role.NormalUserRole) {
			return &exception.ClientException{Msg: "没有投票权限"}
		}

		option, e := dao.GetContentVotingOption(tx, vid)
		if e != nil {
			return &exception.ClientException{Msg: "投票选项错误"}
		}

		if option.Cid != content.ID {
			return &exception.ClientException{Msg: "投票选项不存在"}
		}

		count, e := dao.GetVoteContentCount(tx, uid, cid)
		if e != nil || count > 0 {
			return &exception.ClientException{Msg: "已经投票"}
		}

		e = dao.VoteContent(tx, uid, cid, vid)
		if e != nil {
			return &exception.ClientException{Msg: "投票失败"}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	content, err := GetContent(cid)

	return content, err
}

func CancelVote(uid, cid int64) (*vo.ContentVO, error) {
	err := mysql.GetDB().Transaction(func(tx *gorm.DB) error {
		user, err := dao.GetUserByID(tx, uid)
		if err != nil {
			return &exception.ClientException{Msg: "用户不存在"}
		}

		if !user.Role.Validate(role.NormalUserRole) {
			return &exception.ClientException{Msg: "您还不可以投票哦"}
		}

		err = dao.DeleteVotingInfo(tx, user.ID, cid)
		if err != nil {
			return &exception.ClientException{Msg: "您可能还没有投票哦"}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return GetContent(cid)
}

func DeleteContentVote(uid int64, cid int64) (*vo.ContentVO, error) {
	err := mysql.GetDB().Transaction(func(tx *gorm.DB) error {
		user, err := dao.GetUserByID(tx, uid)
		if err != nil {
			return &exception.ClientException{Msg: "用户不存在"}
		}

		content, err := dao.GetContent(tx, cid)
		if err != nil {
			return &exception.ClientException{Msg: "未查询到该贴子"}
		}

		if !(user.Role.Validate(role.AdminRole, role.SuperUserRole) ||
			(user.Role.Validate(role.NormalUserRole) && content.Uid == user.ID)) {
			return &exception.ClientException{Msg: "您还不可以删除贴子评论哦"}
		}

		if content.Voting == 0 {
			return &exception.ClientException{Msg: "这个帖子还没有开启投票哦"}
		}

		err = dao.UpdateVoteOptionDeleteUid(tx, content.ID, user.ID)
		if err != nil {
			return &exception.ServerException{Msg: "跟新删除信息错误"}
		}

		err = dao.DeleteVote(tx, cid)
		if err != nil {
			return &exception.ServerException{Msg: "删除失败"}
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return GetContent(cid)
}
