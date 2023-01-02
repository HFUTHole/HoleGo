package service

import (
	"hole/pkgs/config"
	"testing"
)

func TestGetReplyPage(t *testing.T) {
	config.InitConfigFileWithTest()
	page, err := GetReplyPage(1, 10, 0)
	if err != nil {
		return
	}

	t.Log(page)
}
