package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"hole/pkgs/config/base"
	"hole/pkgs/models"
	"log"
	"os"
	"time"
)

var cfg *Config
var db *gorm.DB

type Config struct {
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

func InitConfig() {
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

	cfg = &Config{
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

func Init() {
	var err error
	InitConfig()
	db, err = gorm.Open(mysql.Open(cfg.Url), &gorm.Config{})

	if err != nil {
		log.Println("数据库连接失败")
		os.Exit(-1)
	}

	sqlDB, _ := db.DB()

	err = db.AutoMigrate(
		&models.User{},
		&models.TokenInfo{},
		&models.Content{},
		&models.ContentJumpUrl{},
		&models.Tag{},
		&models.ContentTag{},
		&models.ContentImage{},
		&models.VotingOption{},
		&models.VotingInfo{},
		&models.ContentLiked{},
		&models.Reply{},
		&models.AtName{},
	)
	if err != nil {
		fmt.Println("表迁移失败！", err)
	}

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
}

// GetConfig return mysql config
func GetConfig() *Config {
	return cfg
}

func GetDB() *gorm.DB {
	if base.GetMode() == "dev" {
		return db.Debug()
	}
	return db
}
