package dao

import (
	"gorm.io/gorm"
	"hole/pkgs/models"
)

func CreateVote(tx *gorm.DB, cid int64, options []string) error {
	votingOptions := make([]models.VotingOption, len(options))

	for i, option := range options {
		votingOptions[i] = models.VotingOption{
			Cid:  cid,
			Text: option,
		}
	}

	err := tx.Model(&models.VotingOption{}).CreateInBatches(votingOptions, len(votingOptions)).Error
	return err
}

func GetCountContentByCid(tx *gorm.DB, cid int64) (int64, error) {
	var count int64
	err := tx.Model(&models.VotingOption{}).Where("cid = ?", cid).Count(&count).Error
	return count, err
}

func GetContentVoting(tx *gorm.DB, cid int64) ([]models.VotingOption, error) {
	var votingOptions []models.VotingOption
	err := tx.Model(&models.VotingOption{}).Where("cid = ?", cid).Find(&votingOptions).Error
	return votingOptions, err
}

func GetContentVotingOption(tx *gorm.DB, vid int64) (*models.VotingOption, error) {
	var votingOption models.VotingOption
	err := tx.Model(&models.VotingOption{}).Where("id = ?", vid).First(&votingOption).Error
	return &votingOption, err
}

func DeleteVotingInfo(tx *gorm.DB, uid, cid int64) error {
	err := tx.Exec("update voting_options set total = total - 1 where id = (select voting_infos.vid from voting_infos where voting_infos.uid = ? and cid = ?)", uid, cid).Error
	if err != nil {
		return err
	}
	err = tx.Exec("delete from voting_infos where uid = ? and cid = ?", uid, cid).Error
	return err
}

func GetVotingInfo(tx *gorm.DB, uid, cid int64) (*models.VotingInfo, error) {
	var info models.VotingInfo
	err := tx.Model(&models.VotingInfo{}).Where("uid = ? AND cid = ?", uid, cid).First(&info).Error
	return &info, err
}

func VoteContent(tx *gorm.DB, uid int64, cid int64, vid int64) error {
	err := tx.Model(&models.VotingInfo{}).Create(&models.VotingInfo{
		Cid: cid,
		Vid: vid,
		Uid: uid,
	}).Error
	if err != nil {
		return err
	}
	err = tx.Exec("UPDATE `voting_options` SET `total` = `total` + 1 WHERE id = ? AND `voting_options`.`deleted_at` IS NULL", vid).Error
	if err != nil {
		return err
	}
	return nil
}

func GetVoteContentCount(tx *gorm.DB, uid int64, cid int64) (int64, error) {
	var count int64
	err := tx.Model(&models.VotingInfo{}).Where("cid = ? and uid = ?", cid, uid).Count(&count).Error
	return count, err
}
