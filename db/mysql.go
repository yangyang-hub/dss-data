package db

import (
	configs "dss-data/configs"
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Mysql *gorm.DB

// mysql初始化文件
func InitMysql() {
	log.Println("init mysql...")
	db, err := gorm.Open(mysql.Open(configs.Config.MysqlUrl), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})
	if err != nil {
		log.Println("open mysql error:", err)
		panic(err)
	}
	sqldb, _ := db.DB()
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqldb.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqldb.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqldb.SetConnMaxLifetime(time.Hour)
	Mysql = db
	log.Println("init mysql success.")
}

type Writer struct {
}

func (w Writer) Printf(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}
