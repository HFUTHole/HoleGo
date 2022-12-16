package vo

import (
	"encoding/json"
	"gorm.io/gorm"
	"hole/pkgs/models"
	"testing"
)

func TestConvertConvertContentVO(t *testing.T) {
	content := models.Content{
		ID:       1,
		Uid:      20,
		Nick:     "nick",
		Avatar:   "Avatar",
		Like:     30,
		Real:     1,
		Title:    "Title",
		Text:     "Text",
		DeleteAt: gorm.DeletedAt{},
	}

	tags := []models.Tag{{
		ID:   1,
		Text: "hello",
	}}

	urls := []models.ContentJumpUrl{{
		ID:       1,
		Cid:      1,
		Text:     "text",
		JumpCid:  11111,
		DeleteAt: gorm.DeletedAt{},
	}}

	contentVO := ConvertConvertContentVO(&content, tags, urls, nil)
	marshal, err := json.Marshal(contentVO)
	t.Log(string(marshal), err)
}
