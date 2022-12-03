package mysql

import (
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"hole_go/models"
	"hole_go/settings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)


var db gorm.DB

func InitGorm(cfg *settings.MySQLConfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DbName,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		zap.L().Error("connect DB failed", zap.Error(err))
	}

	sqlDB, _ := db.DB()

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		fmt.Println("User表迁移失败！")
	}

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	return
}

func Close() {
	sqlDB, _ := db.DB()
	_ = sqlDB.Close()
}





//var db *sqlx.DB

//func Init(cfg *settings.MySQLConfig) (err error) {
//	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
//		cfg.User,
//		cfg.Password,
//		cfg.Host,
//		cfg.Port,
//		cfg.DbName,
//	)
//	fmt.Println(dsn)
//	// 也可以使用MustConnect连接不成功就panic
//	db, err = sqlx.Connect("mysql", dsn)
//	if err != nil {
//		zap.L().Error("connect DB failed", zap.Error(err))
//		return
//	}
//	db.SetMaxOpenConns(cfg.MaxOpenConns)
//	db.SetMaxIdleConns(cfg.MaxIdleConns)
//	return
//}
//
//func Close() {
//	_ = db.Close()
//}








