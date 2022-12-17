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
