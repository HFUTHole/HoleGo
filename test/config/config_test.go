package config

import (
	"hole/pkgs/config"
	"testing"
)

func TestGetMysqlConfig(t *testing.T) {
	config.InitConfigFileWithTest()
	mysqlConfig := config.GetMysqlConfig()
	t.Log(mysqlConfig)
}
