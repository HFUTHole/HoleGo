package config

import (
	"hole/pkgs/config"
	"hole/pkgs/config/logger"
	"testing"
)

func TestGetLogger(t *testing.T) {
	config.InitConfigFileWithTest()

	logger.InitLogger()

	logger.GetLogger().Info("hello")

	logger.GetSugaredLogger().Infof("hello %s", "debug")
}