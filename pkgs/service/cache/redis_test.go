package cache

import (
	"hole/pkgs/config"
	"hole/pkgs/config/redis"
	"hole/pkgs/models"
	"strconv"
	"testing"
	"time"
)

func TestSetUser(t *testing.T) {
	config.InitConfigFileWithTest()
	redis.InitRedis()

	user := &models.User{
		ID:        0,
		StudentId: 0,
		Username:  "zou yu",
		Password:  "11111",
		Role:      0,
		AuthTime:  time.Time{},
		Sex:       0,
		Avatar:    "1111",
		Email:     "1111",
		About:     "1111",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}

	err := SetUser(strconv.FormatInt(user.ID, 10), user)
	if err != nil {
		t.Error(err)
	}

	getUser, err := GetUser(strconv.FormatInt(user.ID, 10))
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(getUser)
}
