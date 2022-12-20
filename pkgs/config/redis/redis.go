package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"log"
	"os"
)

var cfg *Config

// 声明一个全局的rdb变量
var rdb *redis.Client

type Config struct {
	Host     string `json:"host"`
	Password string `json:"password"`
	Port     int    `json:"port"`
	DB       int    `json:"db"`
	PoolSize int    `json:"pool_size"`
}

func InitConfig() *Config {
	var host = "127.0.0.1"
	var password = ""
	var port = 6379
	var db = 0
	var poolSize = 30

	if value := viper.GetString("redis.host"); value != "" {
		host = value
	}

	if value := viper.GetString("redis.password"); value != "" {
		password = value
	}

	if value := viper.GetInt("redis.port"); value > 0 {
		port = value
	}

	if value := viper.GetInt("redis.db"); value > 0 {
		db = value
	}

	if value := viper.GetInt("redis.poolSize"); value > 0 {
		poolSize = value
	}

	return &Config{
		Host:     host,
		Password: password,
		Port:     port,
		DB:       db,
		PoolSize: poolSize,
	}
}

// Init InitRedis initial redis config
func Init() {
	cfg = InitConfig()
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

// GetRedis return redis client
func GetRedis() *redis.Client {
	return rdb
}

// GetConfig return redis config
func GetConfig() *Config {
	return cfg
}

// CloseRedis close redis
func CloseRedis() {
	_ = rdb.Close()
}
