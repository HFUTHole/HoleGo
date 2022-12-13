package cache

import (
	"encoding/json"
	"hole/src/config/redis"
	"hole/src/exception"
	"hole/src/models"
	"time"
)

const Prefix = "user:"

func SetUser(uid string, user *models.User) error {
	client := redis.GetRedis()

	marshal, err := json.Marshal(user)

	if err != nil {
		return &exception.ServerException{
			Msg: "缓存异常",
		}
	}
	key := Prefix + uid
	err = client.Set(key, string(marshal), time.Minute*2).Err()

	if err != nil {
		return &exception.ServerException{
			Msg: "缓存存入异常",
		}
	}

	return nil
}

func GetUser(uid string) (*models.User, error) {

	client := redis.GetRedis()
	stringCmd := client.Get(Prefix + uid)

	if stringCmd.Err() != nil {
		return nil, &exception.BusinessException{Msg: "获取缓存失败 uid: " + uid}
	}

	client.Expire(Prefix+uid, time.Minute*2)

	var user = new(models.User)
	err := json.Unmarshal([]byte(stringCmd.Val()), user)
	if err != nil {
		return nil, &exception.BusinessException{Msg: "获取缓存序列化 uid: " + uid}
	}

	return user, nil
}
