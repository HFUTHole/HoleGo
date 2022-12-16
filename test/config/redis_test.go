package config

import (
	"hole/pkgs/config"
	"hole/pkgs/config/redis"
	"testing"
	"time"
)

func TestGetRedis(t *testing.T) {
	config.InitConfigFileWithTest()
	redis.InitRedis()

	err := redis.GetRedis().Set("name", "zou yu", time.Duration(time.Second*100)).Err()
	if err != nil {
		t.Error(err)
		return
	}

	val := redis.GetRedis().Get("name").Val()
	t.Log(val)
}
