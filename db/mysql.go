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
	param_s := fmt.Sprintf(
		"%v:%v@tcp(%v:%v)/%v?parseTime=True&loc=Local",
		configs.Config.MysqlUsername,
		configs.Config.Mysqlpassword,
		configs.Config.MysqlHost,
		configs.Config.MysqlPort,
		configs.Config.MysqlDatabase,
	)
	log.Println("mysql param:", param_s)
	db, err := gorm.Open(mysql.Open(param_s), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})
	if err != nil {
		log.Println("open mysql error:", err)
		panic(err)
	}
	sqldb, err := db.DB()
	sqlDB := *sqldb
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(20)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
	Mysql = db
	log.Println("init mysql end.")
}

type Writer struct {
}

func (w Writer) Printf(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}
