package dao

import (
	"errors"
	"gorm.io/gorm"
	"hole/pkgs/models"
)

func CreateUser(tx *gorm.DB, user *models.User) error {
	err := tx.Create(user).Error
	return err
}

func GetUserByStudentID(tx *gorm.DB, studentId int64) (*models.User, error) {
	var user models.User

	err := tx.Model(&models.User{}).Where("student_id = ?", studentId).First(&user).Error
	if err != nil {
		return nil, errors.New("未查询到用户")
	}
	return &user, nil
}

func GetUserByID(tx *gorm.DB, id int64) (*models.User, error) {
	var user models.User

	err := tx.Model(&models.User{}).Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, errors.New("未查询到用户")
	}
	return &user, nil
}
