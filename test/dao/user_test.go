package dao

import (
	"gorm.io/gorm"
	"hole/pkgs/config"
	"hole/pkgs/config/mysql"
	"hole/pkgs/dao"
	"hole/pkgs/models"
	"hole/pkgs/models/role"
	"testing"
	"time"
)

func GetDB() *gorm.DB {
	config.InitConfigFileWithTest()
	mysql.Init()
	return mysql.GetDB()
}

func TestCreateUser(t *testing.T) {
	db := GetDB()

	err := db.Debug().Transaction(func(tx *gorm.DB) error {
		err := dao.CreateUser(tx, &models.User{
			StudentId: 2020218081,
			Username:  "AdminRole",
			Password:  "1002",
			Role:      role.CreateRole(role.AdminRole),
			AuthTime:  time.Now(),
			Sex:       0,
			Avatar:    "icon.png",
			Email:     "user.zouyu@foxmail.com",
			About:     "zou yu",
		})
		if err != nil {
			t.Fatal(err)
			return err
		}
		user, err := dao.GetUserByStudentID(tx, 2020218081)

		if err != nil {
			t.Fatal(err)
			return err
		}
		t.Log(user)
		return nil
	})

	if err != nil {
		t.Fatal(err)
		return
	}

}

func TestGetUserByStudentID(t *testing.T) {
	db := GetDB()

	err := db.Debug().Transaction(func(tx *gorm.DB) error {
		user, err := dao.GetUserByStudentID(tx, 2020218081)

		if err != nil {
			t.Fatal(err)
			return err
		}
		t.Log(user)

		return nil
	})
	if err != nil {
		t.Error(err)
		return
	}
}

func TestGetUserByID(t *testing.T) {
	db := GetDB()

	err := db.Debug().Transaction(func(tx *gorm.DB) error {
		user, err := dao.GetUserByID(tx, 1)
		if err != nil {
			t.Fatal(err)
			return err
		}
		t.Log(user)
		return nil
	})

	if err != nil {
		return
	}
}
