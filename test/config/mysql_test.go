package config

import (
	"gorm.io/gorm"
	"hole/src/config"
	"hole/src/config/mysql"
	"hole/src/models"
	"testing"
	"time"
)

func TestGetDB(t *testing.T) {
	config.InitConfigFileWithTest()
	mysql.InitMysql()

	getDB := mysql.GetDB()

	t.Log(getDB)
}

func TestCreate(t *testing.T) {
	config.InitConfigFileWithTest()
	mysql.InitMysql()
	db := mysql.GetDB()

	err := db.Debug().Transaction(func(tx *gorm.DB) error {
		tx.Begin()
		//err := tx.AutoMigrate(&User{})
		//if err != nil {
		//	return err
		//}

		tx.Model(&models.User{}).Create(&models.User{
			StudentId: 0,
			Username:  "",
			Password:  "",
			Role:      0,
			AuthTime:  time.Now(),
			Sex:       0,
			Avatar:    "",
			Email:     "",
			About:     "",
		})
		return nil
	})

	if err != nil {
		return
	}
}
