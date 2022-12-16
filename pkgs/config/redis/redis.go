package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"hole/pkgs/config"
	"log"
	"os"
)

// 声明一个全局的rdb变量
var rdb *redis.Client

// InitRedis 初始化连接
func InitRedis() {
	cfg := config.GetRedisConfig()
	rdb = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			cfg.Host,
			cfg.Port,
		),
		Password: cfg.Password, // no password set
		DB:       cfg.DB,       // use default DB
		PoolSize: cfg.PoolSize,
	})
	_, err := rdb.Ping().Result()
	if err != nil {
		log.Fatalln(err)
		os.Exit(-1)
	}
}

func GetRedis() *redis.Client {
	return rdb
}

func CloseRedis() {
	_ = rdb.Close()
}
