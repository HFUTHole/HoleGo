package dao

import (
	"gorm.io/gorm"
	"hole/pkgs/config"
	"hole/pkgs/config/mysql"
	"hole/pkgs/models"
	"testing"
	"time"
)

func TestCreateReply(t *testing.T) {
	config.InitConfigFileWithTest()
	db := mysql.GetDB()

	err := CreateReply(db, &models.Reply{
		Root:      -1,
		Parent:    -1,
		Cid:       1,
		Uid:       2,
		Real:      0,
		Nick:      "zou yu",
		Avatar:    "",
		Message:   "hi zou yu",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeleteAt:  gorm.DeletedAt{},
	})

	if err != nil {
		t.Error(err)
		return
	}

}

func TestGetRootReply(t *testing.T) {
	config.InitConfigFileWithTest()
	db := mysql.GetDB()
	replies, err := GetRootReplies(db, 1, 10)

	if err != nil {
		t.Fatal(err)
	}

	t.Log(replies)
}

func TestGetChildren(t *testing.T) {
	config.InitConfigFileWithTest()
	db := mysql.GetDB()

	children, err := GetChildren(db, 1, 1)

	if err != nil {
		t.Fatal(err)
	}
	t.Log(children)
}
