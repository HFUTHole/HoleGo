package dao

import (
	"gorm.io/gorm"
	"hole/pkgs/models"
)

func CreateReply(tx *gorm.DB, reply *models.Reply) error {
	return tx.Model(&models.Reply{}).Create(reply).Error
}

func GetParentReply(tx *gorm.DB, cid, parent int64) (*models.Reply, error) {
	var reply models.Reply
	err := tx.Model(&models.Reply{}).Where("cid = ? and id = ?", cid, parent).First(&reply).Error
	return &reply, err
}

func GetRootReply(tx *gorm.DB, cid, root int64) (*models.Reply, error) {
	var reply models.Reply
	err := tx.Model(&models.Reply{}).Where("root = -1 and cid = ? and id = ?", cid, root).First(&reply).Error
	return &reply, err
}

func GetRootReplies(tx *gorm.DB, cid int64, pageSize int) ([]models.Reply, error) {
	var replies []models.Reply
	err := tx.Model(&models.Reply{}).Where("cid = ? and root = -1", cid).Order("id desc").Limit(pageSize).Find(&replies).Error
	return replies, err
}

func GetRootReplyNext(tx *gorm.DB, cid int64, pageSize int, maxId int64) ([]models.Reply, error) {
	var replies []models.Reply
	err := tx.Model(&models.Reply{}).Where("cid = ? and root = -1 and id < ?", cid, maxId).Order("id desc").Limit(pageSize).Find(&replies).Error
	return replies, err
}

func GetChildren(tx *gorm.DB, cid int64, rootId int64) ([]models.Reply, error) {
	var replies []models.Reply
	err := tx.Model(&models.Reply{}).Where("cid = ? and root = ?", cid, rootId).Order("id desc").Find(&replies).Error
	return replies, err
}

func CreateAtName(tx *gorm.DB, at *models.AtName) error {
	err := tx.Model(&models.AtName{}).Create(at).Error
	return err
}

func GetAtNameByRid(tx *gorm.DB, rid int64) ([]models.AtName, error) {
	var list []models.AtName
	err := tx.Model(&models.AtName{}).Where("reply_id = ?", rid).Find(&list).Error
	return list, err
}
