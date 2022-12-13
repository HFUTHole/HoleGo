package config

import (
	"hole/src/config"
	"hole/src/config/logger"
	"testing"
)

func TestGetLogger(t *testing.T) {
	config.InitConfigFileWithTest()

	logger.InitLogger()

	logger.GetLogger().Info("hello")

	logger.GetSugaredLogger().Infof("hello %s", "debug")
}
