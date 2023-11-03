package cache

import (
	"ace/model"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var DB *gorm.DB

func InitMysql(conf model.Mysql) {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True", conf.Username, conf.Password, conf.Host, conf.Port, conf.DB)
	config := gorm.Config{
		Logger: logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,         // Disable color
		}),
		PrepareStmt: true,
	}
	db, err := gorm.Open(mysql.Open(dsn), &config)
	if err != nil {
		fmt.Printf("Open mysql error: %v \n", err.Error())
		os.Exit(1)
	}

	sqlDB, err := db.DB()
	if err != nil {
		fmt.Printf("DealWith mysql error: %v \n", err.Error())
		os.Exit(1)
	}

	sqlDB.SetMaxIdleConns(50)
	sqlDB.SetMaxOpenConns(200)

	DB = db

	// 自动迁移
	migration()
}
