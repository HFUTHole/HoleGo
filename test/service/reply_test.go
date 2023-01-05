package service

import (
	"hole/pkgs/config"
	"hole/pkgs/service"
	"testing"
)

func TestGetReplyPage(t *testing.T) {
	config.InitConfigFileWithTest()
	page, err := service.GetReplyPage(1, 10, 0)
	if err != nil {
		return
	}

	t.Log(page)
}

func TestGetReply(t *testing.T) {
	config.InitConfigFileWithTest()
	reply, err := service.GetReply(1605499471414693888, 1610634456618504192)
	if err != nil {
		t.Error(err)
	}

	t.Log(reply)
}
