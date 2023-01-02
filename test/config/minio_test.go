package config

import (
	"hole/pkgs/config"
	"hole/pkgs/config/fileservice"
	"testing"
)

func TestExists(t *testing.T) {
	config.InitConfigFileWithTest()
	exists := fileservice.Exists(fileservice.TempBucket, "1605207910286102528")

	t.Log(exists)
}
