package config

import (
	"hole/src/config"
	"testing"
)

func TestGetMysqlConfig(t *testing.T) {
	config.InitConfigFileWithTest()
	mysqlConfig := config.GetMysqlConfig()
	t.Log(mysqlConfig)
}
