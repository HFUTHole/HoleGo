package dao

import (
	"gorm.io/gorm"
	"hole/src/models"
)

func CreateToken(tx *gorm.DB, login *models.TokenInfo) error {
	err := tx.Model(&models.TokenInfo{}).Create(login).Error
	return err
}

func GetAllToken(tx *gorm.DB, audience int64) []models.TokenInfo {
	var login []models.TokenInfo
	err := tx.Model(&models.TokenInfo{}).Where("audience = ?", audience).Find(&login).Error
	if err != nil {
		return nil
	}

	return login
}

func GetTokenByID(tx *gorm.DB, id int64) (*models.TokenInfo, error) {
	var login models.TokenInfo
	err := tx.Model(models.TokenInfo{}).Where("id = ?", id).First(&login).Error
	return &login, err
}
