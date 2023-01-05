package dao

import (
	"gorm.io/gorm"
	"hole/pkgs/config"
	"hole/pkgs/config/mysql"
	"hole/pkgs/dao"
	"hole/pkgs/models"
	"testing"
	"time"
)

func TestCreateReply(t *testing.T) {
	config.InitConfigFileWithTest()
	db := mysql.GetDB()

	err := dao.CreateReply(db, &models.Reply{
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
	replies, err := dao.GetRootReplies(db, 1, 10)

	if err != nil {
		t.Fatal(err)
	}

	t.Log(replies)
}

func TestGetChildren(t *testing.T) {
	config.InitConfigFileWithTest()
	db := mysql.GetDB()

	children, err := dao.GetChildren(db, 1, 1)

	if err != nil {
		t.Fatal(err)
	}
	t.Log(children)
}

func TestDeleteReply(t *testing.T) {
	config.InitConfigFileWithTest()

	err := dao.DeleteReply(mysql.GetDB(), 1609528753388523520)

	t.Log(err)
}

func TestGetRepliesIdByParent(t *testing.T) {
	config.InitConfigFileWithTest()

	parent, err := dao.GetRepliesIdByParent(mysql.GetDB(), []int64{1610612779121643520})

	if err != nil {
		t.Fatal(err)
	}

	t.Log(parent)
	err = dao.DeleteReplies(mysql.GetDB(), []int64{1610612779121643520})
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetReplyByID(t *testing.T) {
	config.InitConfigFileWithTest()
	id, err := dao.GetReplyByID(mysql.GetDB(), 1)
	if err != nil {
		t.Error(err)
	}
	t.Log(id)

}
