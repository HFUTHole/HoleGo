package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"hole/pkgs/config"
	"hole/pkgs/models"
	"log"
	"os"
	"time"
)

var db *gorm.DB

func InitMysql() {
	var err error
	mysqlConfig := config.GetMysqlConfig()
	db, err = gorm.Open(mysql.Open(mysqlConfig.Url), &gorm.Config{})

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
	)
	if err != nil {
		fmt.Println("表迁移失败！", err)
	}

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(mysqlConfig.MaxIdleConns)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(mysqlConfig.MaxOpenConns)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
}

func GetDB() *gorm.DB {
	if config.GetMode() == "dev" {
		return db.Debug()
	}
	return db
}
