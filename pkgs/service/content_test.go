package service

import (
	"encoding/json"
	"hole/pkgs/config"
	"hole/pkgs/config/logger"
	"hole/pkgs/config/mysql"
	"testing"
)

func TestCreateContent(t *testing.T) {
	config.InitConfigFileWithTest()
	config.InitUtils()
	logger.InitLogger()
	mysql.InitMysql()

	content, err := CreateContent(
		1,
		"title",
		"# zou yu",
		[]string{"test 1"},
		[]string{"https://hello.png"},
		true,
	)
	if err != nil {
		t.Error(err)
	}

	t.Log(content)
}

func TestSearchMessageJumpUrls(t *testing.T) {
	config.InitConfigFileWithTest()
	config.InitUtils()
	//logger.InitLogger()
	mysql.InitMysql()
	urls, err := SearchMessageJumpUrls("#1603047268133376000 #1603047268133376000 hello", 1)
	if err != nil {
		t.Error(err)
	}

	t.Log(urls)
}

func TestGetContent(t *testing.T) {
	config.InitConfigFileWithTest()
	config.InitUtils()
	//logger.InitLogger()
	mysql.InitMysql()
	content, err := GetContent(1603642730045509632)
	if err != nil {
		t.Error(err)
	}
	marshal, err := json.Marshal(content)
	t.Log(string(marshal))
}

func TestGetContentPage(t *testing.T) {
	config.InitConfigFileWithTest()
	config.InitUtils()
	//logger.InitLogger()
	mysql.InitMysql()

	page, err := GetContentPage(10)
	if err != nil {
		t.Error(err)
	}

	t.Log(page)
}

func TestGetContentNextPage(t *testing.T) {
	config.InitConfigFileWithTest()
	config.InitUtils()
	//logger.InitLogger()
	mysql.InitMysql()

	page, err := GetContentNextPage(1603690087504154624, 10)
	if err != nil {
		t.Error(err)
	}
	// 1603690067446992896
	// 1603672757168508928

	t.Log(page)
}
