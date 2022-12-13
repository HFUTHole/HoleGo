package config

import (
	"github.com/spf13/viper"
	"go.uber.org/zap/zapcore"
	"hole/src/pkg/utils"
	"log"
	"os"
)

func InitConfigFile() {
	viper.SetConfigType("yaml")
	viper.SetConfigFile("conf/config.yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()

	if err != nil {
		log.Printf("configuration file read error: %v\n", err)
	}
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

func InitUtils() {
	utils.InitSnowflakeNode()
}

type MySQLConfig struct {
	Url          string
	Username     string
	Password     string
	Host         string
	Port         string
	Db           string
	Param        string
	MaxIdleConns int
	MaxOpenConns int
}

func GetMysqlConfig() *MySQLConfig {
	var url string
	var username string
	var password string
	var host string
	var port string
	var db string
	var param string
	var maxIdleConns = 10
	var MaxOpenConns = 50

	url = os.Getenv("MYSQL_URL")
	if url == "" {
		url = viper.GetString("mysql.url")
	}

	if url == "" {
		username = viper.GetString("mysql.username")
		password = viper.GetString("mysql.password")
		host = viper.GetString("mysql.host")
		if host == "" {
			host = "127.0.0.1"
		}

		port = viper.GetString("mysql.port")
		if port == "" {
			port = "3306"
		}

		db = viper.GetString("mysql.db")
		if db == "" {
			db = "test"
		}

		param = viper.GetString("mysql.param")

		url = username + ":" + password + "@tcp(" + host + ":" + port + ")/" + db
		if param != "" {
			url += "?" + param
		}
	} else {
		right := 0
		l := len(url)
		for left := 0; right < l; right++ {
			if url[right] == ':' {
				username = url[left:right]
				right += 1
				break
			}
		}

		for left := right; right < l; right++ {
			if url[right:right+5] == "@tcp(" {
				password = url[left:right]
				right += 5
				break
			}
		}

		for left := right; right < l; right++ {
			if url[right] == ':' {
				host = url[left:right]
				right += 1
				break
			}
		}

		for left := right; right < l; right++ {
			if url[right:right+2] == ")/" {
				port = url[left:right]
				right += 2
				break
			}
		}

		for left := right; right < l; right++ {
			if url[right] == '?' {
				db = url[left:right]
				right += 1
				break
			}
		}

		param = url[right:]
	}

	if size := viper.GetInt("mysql.max_idle_conns"); size > 0 {
		maxIdleConns = size
	}

	if size := viper.GetInt("mysql.max_open_conns"); size > 0 {
		MaxOpenConns = size
	}

	return &MySQLConfig{
		Url:          url,
		Username:     username,
		Password:     password,
		Host:         host,
		Port:         port,
		Db:           db,
		Param:        param,
		MaxIdleConns: maxIdleConns,
		MaxOpenConns: MaxOpenConns,
	}
}

type LogConfig struct {
	Level      zapcore.Level `json:"level"`
	Filename   string        `json:"filename"`
	MaxSize    int           `json:"max_size"`
	MaxAge     int           `json:"max_age"`
	MaxBackups int           `json:"max_backups"`
}

func GetLoggerConfig() *LogConfig {
	var level = zapcore.InfoLevel
	var filename = ""
	var maxSize = 300
	var maxAge = 30
	var maxBackups = 7

	filename = viper.GetString("log.filename")
	if ls := viper.GetString("log.level"); ls != "" {
		var l = zapcore.InfoLevel
		err := l.UnmarshalText([]byte(ls))
		if err == nil {
			level = l
		}
	}

	if ms := viper.GetInt("log.max_size"); ms > 0 {
		maxSize = ms
	}

	if ma := viper.GetInt("log.max_age"); ma > 0 {
		maxAge = ma
	}

	if mb := viper.GetInt("log.max_age"); mb > 0 {
		maxBackups = mb
	}

	return &LogConfig{
		Level:      level,
		Filename:   filename,
		MaxSize:    maxSize,
		MaxAge:     maxAge,
		MaxBackups: maxBackups,
	}
}

type RedisConfig struct {
	Host     string `json:"host"`
	Password string `json:"password"`
	Port     int    `json:"port"`
	DB       int    `json:"db"`
	PoolSize int    `json:"pool_size"`
}

func GetRedisConfig() *RedisConfig {
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

	return &RedisConfig{
		Host:     host,
		Password: password,
		Port:     port,
		DB:       db,
		PoolSize: poolSize,
	}
}

func GetMode() string {
	mode := viper.GetString("mode")
	if mode == "" {
		mode = "dev"
	}
	return mode
}

func GetPort() int {
	port := viper.GetInt("port")
	if port <= 0 {
		port = 8080
	}
	return port
}

func GetVersion() string {
	version := viper.GetString("version")
	if version == "" {
		version = "未设置版本号"
	}
	return version
}
