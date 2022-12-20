package config

import (
	"github.com/spf13/viper"
	"hole/pkgs/config/base"
	"hole/pkgs/config/fileservice"
	"hole/pkgs/config/logger"
	"hole/pkgs/config/mysql"
	"hole/pkgs/config/redis"
	"log"
)

// Init Read configuration files and assemble basic components, such as mysql, redis, logger, minio,...
func Init() {
	viper.SetConfigType("yaml")
	viper.SetConfigFile("conf/config.yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()

	if err != nil {
		log.Printf("configuration file read error: %v\n", err)
	}
	InitComponents()
}

func InitConfigFileWithTest() {
	viper.SetConfigType("yaml")
	viper.SetConfigFile("C:\\Users\\zouyu\\Desktop\\Porject\\tree-hole\\HoleGo\\conf\\config.yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()

	if err != nil {
		log.Printf("configuration file read error: %v\n", err)
	}
}

func InitComponents() {
	base.Init()
	logger.Init()
	mysql.Init()
	redis.Init()
	fileservice.Init()
}
