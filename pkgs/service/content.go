package service

import (
	"github.com/bwmarrin/snowflake"
	"gorm.io/gorm"
	"hole/pkgs/common/utils"
	"hole/pkgs/config/fileservice"
	"hole/pkgs/config/mysql"
	"hole/pkgs/dao"
	"hole/pkgs/exception"
	"hole/pkgs/models"
	"hole/pkgs/models/role"
	"hole/pkgs/models/vo"
	"strconv"
	"time"
)

var node, _ = snowflake.NewNode(1)

func NextID() int64 {
	return int64(node.Generate())
}

func CreateContent(uid int64, title string, message string, tags []string, urls []string, real bool) (*vo.ContentVO, error) {
	title = utils.Scape(title)
	if len(title) > 32 {
		return nil, &exception.ClientException{Msg: "帖子标题超出规定长度"}
	}

	message = utils.Scape(message)
	if len(message) >= 2048 {
		return nil, &exception.ClientException{Msg: "帖子内容超出规定长度"}
	}

	tags = utils.ScapeSlice(tags)
	if !utils.SliceElementMaxLength(tags, 32) {
		return nil, &exception.ClientException{Msg: "帖子标签超出规定长度"}
	}

	if !utils.SliceElementMaxLength(urls, 32) {
		return nil, &exception.ClientException{Msg: "帖子照片路径超出规定长度"}
	}

	cid := NextID()

	db := mysql.GetDB()
	err := db.Transaction(func(tx *gorm.DB) error {
		user, err := dao.GetUserByID(tx, uid)
		if err != nil {
			return &exception.ClientException{Msg: "未查询到该用户"}
		}

		if user.Role.Validate(role.NormalUserRole, role.AdminRole, role.SuperUserRole) {
			return &exception.ClientException{Msg: "没有写帖子权限"}
		}

		count, err := dao.GetContentOneDayCount(tx, time.Now())
		if err != nil {
			return &exception.ServerException{Msg: "帖子创建失败"}
		}

		if count >= 10 {
			return &exception.ClientException{Msg: "今日不可以再发了哦"}
		}

		content := &models.Content{
			ID:      cid,
			Uid:     user.ID,
			Nick:    user.Username,
			Avatar:  user.Avatar,
			Like:    0,
			Real:    utils.BoolToInt(real),
			Title:   title,
			Text:    message,
			EndTime: time.Now(),
		}
		err = dao.CreateContent(tx, content)
		if err != nil {
			return &exception.ServerException{Msg: "帖子内容创建失败"}
		}

		if tags != nil && len(tags) > 0 {
			err = dao.CreateAndLinkTags(tx, cid, tags)
			if err != nil {
				return &exception.ServerException{Msg: "帖子标签创建失败"}
			}
		}

		jumpUrls, err := SearchMessageJumpUrls(message, cid)
		if err != nil {
			return err
		}

		if jumpUrls != nil && len(jumpUrls) > 0 {
			err = dao.CreateContentJumpUrls(tx, jumpUrls)
			if err != nil {
				return &exception.ServerException{Msg: "帖子引用文章创建失败"}
			}
		}

		if urls != nil && len(urls) > 0 {
			err = dao.CreateContentImages(tx, cid, urls)
			if err != nil {
				return &exception.ServerException{Msg: "帖子照片创建失败"}
			}
		}

		for _, id := range urls {
			err := fileservice.CopyFileToContent(id)
			if err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	content, err := GetContent(cid)
	return content, err
}

func GetContent(cid int64) (*vo.ContentVO, error) {
	var contentVO *vo.ContentVO

	db := mysql.GetDB()
	err := db.Transaction(func(tx *gorm.DB) error {
		content, err := dao.GetContentByID(tx, cid)

		if err != nil {
			return &exception.BusinessException{Msg: "未查询到该帖子"}
		}

		tags, err := dao.GetContentTags(tx, cid)

		if err != nil {
			return &exception.BusinessException{Msg: "帖子标签查询错误"}
		}

		urls, err := dao.GetContentJumpUrls(tx, cid)
		if err != nil {
			return &exception.BusinessException{Msg: "帖子跳转路径查询错误"}
		}

		images, err := dao.GetContentImages(tx, cid)
		if err != nil {
			return &exception.BusinessException{Msg: "帖子照片查询错误"}
		}

		for i := range images {
			images[i].URL = utils.ImageIdToUrl(images[i].URL, fileservice.ContentBucket)
		}

		contentVO = vo.ConvertConvertContentVO(content, tags, urls, images)
		voting, err := GetContentVoting(cid)
		if err != nil {
			return &exception.BusinessException{Msg: "帖子投票查询错误"}
		}
		contentVO.Voting = voting

		return nil
	})
	if err != nil {
		return nil, err
	}

	return contentVO, nil
}

func SearchMessageJumpUrls(message string, parent int64) ([]models.ContentJumpUrl, error) {
	urls := make([]models.ContentJumpUrl, 5)

	length := len(message)
	db := mysql.GetDB()

	idx := 0
	for i := 0; i < length; i++ {
		if message[i] == '#' {
			if idx >= 5 {
				return nil, &exception.ClientException{Msg: "最多引用 5 篇帖子"}
			}
			j := i + 1
			for ; j < length && message[j] != ' '; j++ {
			}
			text := message[i+1 : j]
			cid, err := strconv.ParseInt(text, 10, 64)
			if err != nil {
				return nil, &exception.ClientException{Msg: "错误的引用: #" + text}
			}
			content, err := dao.GetContentByID(db, cid)
			if err != nil {
				return nil, &exception.ClientException{Msg: "错误的引用: #" + text}
			}

			urls[idx] = models.ContentJumpUrl{
				Cid:     parent,
				Text:    text,
				JumpCid: content.ID,
			}
			i = j
			idx++
		}
	}

	return urls[0:idx], nil
}

func GetContentPage(pageSize int) (*vo.ContentPage, error) {
	if pageSize > 20 {
		pageSize = 20
	}

	db := mysql.GetDB()
	page, err := dao.GetContentPage(db, pageSize)
	if err != nil {
		return nil, &exception.ServerException{Msg: "查询错误"}
	}

	var maxID = int64(-1)
	if len(page) > 0 {
		maxID = page[len(page)-1].ID
	}

	return &vo.ContentPage{
		MaxId: maxID,
		List:  vo.ConvertContentVOList(page),
	}, nil
}

func GetContentNextPage(maxId int64, pageSize int) (*vo.ContentPage, error) {
	if pageSize > 20 {
		pageSize = 20
	}

	db := mysql.GetDB()
	page, err := dao.GetContentNextPage(db, maxId, pageSize)
	if err != nil {
		return nil, &exception.ServerException{Msg: "查询错误"}
	}

	var maxID = int64(-1)
	if len(page) > 0 {
		maxID = page[len(page)-1].ID
	}

	return &vo.ContentPage{
		MaxId: maxID,
		List:  vo.ConvertContentVOList(page),
	}, nil
}

func DeleteContent(uid int64, cid int64) error {
	db := mysql.GetDB()

	err := db.Transaction(func(tx *gorm.DB) error {
		user, err := dao.GetUserByID(tx, uid)
		if err != nil {
			return &exception.ClientException{Msg: "未查询到该用户"}
		}

		if user.Role.Validate(role.NormalUserRole, role.AdminRole, role.SuperUserRole) {
			return &exception.ClientException{Msg: "没有写帖子权限"}
		}

		content, err := dao.GetContent(tx, cid)
		if err != nil {
			return &exception.ServerException{Msg: "帖子查询错误"}
		}

		if content.Uid != user.ID {
			return &exception.ClientException{Msg: "这不是你的帖子哦"}
		}
		if content.DeleteAt.Valid == true {
			return &exception.ClientException{Msg: "帖子已经删除"}
		}

		err = dao.DeleteContentByID(tx, cid)
		if err != nil {
			return &exception.ClientException{Msg: "删除错误"}
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func CreateLiked(uid int64, cid int64) (*vo.ContentVO, error) {
	err := mysql.GetDB().Transaction(func(tx *gorm.DB) error {
		user, err := dao.GetUserByID(tx, uid)
		if err != nil {
			return &exception.ClientException{Msg: "未查询到该用户"}
		}

		if user.Role.Validate(role.NormalUserRole, role.AdminRole, role.SuperUserRole) {
			return &exception.ClientException{Msg: "没有点赞权限"}
		}

		err = dao.CreateLiked(tx, uid, cid)
		if err != nil {
			return &exception.ClientException{Msg: "已点赞"}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return GetContent(cid)
}

func CancelLiked(uid int64, cid int64) (*vo.ContentVO, error) {
	err := mysql.GetDB().Transaction(func(tx *gorm.DB) error {
		user, err := dao.GetUserByID(tx, uid)
		if err != nil {
			return &exception.ClientException{Msg: "未查询到该用户"}
		}

		if user.Role.Validate(role.NormalUserRole, role.AdminRole, role.SuperUserRole) {
			return &exception.ClientException{Msg: "没有点赞权限"}
		}

		err = dao.CancelLiked(tx, uid, cid)
		if err != nil {
			return &exception.ClientException{Msg: "你可能还没有点赞哦"}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return GetContent(cid)
}
