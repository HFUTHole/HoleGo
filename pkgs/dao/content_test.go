package dao

import (
	"gorm.io/gorm"
	"hole/pkgs/models"
	"testing"
	"time"
)

func TestCreateContent(t *testing.T) {
	db := GetDB()

	c := &models.Content{
		ID:       0,
		Uid:      0,
		Nick:     "zou yu",
		Avatar:   "zou yu",
		Like:     0,
		Real:     0,
		Title:    "zou yu",
		Text:     "zou yu",
		DeleteAt: gorm.DeletedAt{},
	}
	err := CreateContent(db, c)
	if err != nil {
		t.Error(err)
	}
}

func TestGetContentFormID(t *testing.T) {
	db := GetDB()

	content, err := GetContentFormID(db, 160304679018404659)
	if err != nil {
		t.Error(err)
	}
	t.Log(content)
}

func TestGetContentList(t *testing.T) {
	db := GetDB()
	list, err := GetContentNextPage(db, 4, 2)
	if err != nil {
		t.Error(err)
	}
	t.Log(list)
}

func TestGetContentPage(t *testing.T) {
	db := GetDB()
	list, err := GetContentPage(db, 4)
	if err != nil {
		t.Error(err)
	}
	t.Log(list)
}

func TestGetContentJumpUrl(t *testing.T) {
	db := GetDB()
	list, err := GetContentJumpUrls(db, 1)
	if err != nil {
		t.Error(err)
	}
	t.Log(list)
}

func TestCreateContentJumpUrl(t *testing.T) {
	db := GetDB()
	c := &models.ContentJumpUrl{
		Cid:      1,
		Text:     "Test ",
		JumpCid:  122,
		DeleteAt: gorm.DeletedAt{},
	}

	err := CreateContentJumpUrl(db, c)
	if err != nil {
		t.Error(err)
	}
}

func TestCreateContentJumpUrls(t *testing.T) {
	db := GetDB()
	var list = make([]models.ContentJumpUrl, 2)
	list[0] = models.ContentJumpUrl{
		Cid:      1,
		Text:     "Test ",
		JumpCid:  122,
		DeleteAt: gorm.DeletedAt{},
	}
	list[1] = models.ContentJumpUrl{
		Cid:      1,
		Text:     "Test 2",
		JumpCid:  122,
		DeleteAt: gorm.DeletedAt{},
	}

	err := CreateContentJumpUrls(db, list)
	if err != nil {
		t.Error(err)
	}
}

func TestCreateTag(t *testing.T) {
	db := GetDB()
	tag := &models.Tag{
		Text: "test 2",
	}
	err := CreateTag(db, tag)
	if err != nil {
		t.Error(err)
	}
}

func TestGetTag(t *testing.T) {
	db := GetDB()
	tag, err := GetTag(db, 1)
	if err != nil {
		t.Error(err)
	}
	t.Log(tag)
}

func TestGetTagByText(t *testing.T) {
	db := GetDB()
	tag, err := GetTagByText(db, "test")

	if err != nil {
		t.Error(err)
	}
	t.Log(tag)
}

func TestGetContentTags(t *testing.T) {
	db := GetDB()
	tags, err := GetContentTags(db, 1)
	if err != nil {
		t.Error(err)
	}
	t.Log(tags)
}

func TestLinkContentTag(t *testing.T) {
	db := GetDB()
	link := &models.ContentTag{
		Cid: 1,
		Tid: 2,
	}
	err := LinkContentTag(db, link)
	if err != nil {
		t.Error(err)
	}
}

func TestCreateAndLinkTags(t *testing.T) {
	db := GetDB()
	var tags = make([]string, 2)
	tags[0] = "test"
	tags[1] = "test 2"

	err := CreateAndLinkTags(db, 2, tags)
	if err != nil {
		t.Error(err)
	}
}

func TestCreateContentImages(t *testing.T) {
	db := GetDB()
	err := CreateContentImages(db, 1, []string{"111.png", "2.png"})
	if err != nil {
		t.Error(err)
		return
	}
}

func TestGetContentImages(t *testing.T) {
	db := GetDB()
	images, err := GetContentImages(db, 1)
	if err != nil {
		t.Error(err)
	}
	t.Log(images)

}

func TestGetContentByUid(t *testing.T) {
	db := GetDB()
	uid, err := GetContentOffset10ByUid(db, 1)
	if err != nil {
		t.Error(err)
	}
	t.Log(time.Now().Sub(uid.CreatedAt) < time.Hour*24)
}

func TestGetContentOneDay(t *testing.T) {
	db := GetDB()

	day, err := GetContentOneDay(db, time.Now())
	if err != nil {
		t.Error(err)
	}
	t.Log(len(day))

}

func TestGetContentOneDayCount(t *testing.T) {
	db := GetDB()

	day, err := GetContentOneDayCount(db, time.Now())
	if err != nil {
		t.Error(err)
	}
	t.Log(day)
}

func TestSetContentVoteEndTime(t *testing.T) {
	db := GetDB()

	err := SetContentVoteEndTime(db, 1603642475187015680, 1, time.Now().Add(time.Hour*3))
	if err != nil {
		t.Fatal(err)
		return
	}
}
