package main

import (
	"crypto/md5"
	"encoding/hex"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"hcshop_srvs/user_srv/model"
	"io"
	"log"
	"os"
	"time"
)

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/hcshop_user_srv?charset=utf8mb4&parseTime=True&loc=Local"
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // 禁用彩色打印
		},
	)

	// 全局模式
	//NamingStrategy和Tablename不能同时配置，
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&model.User{})
}


func genMd5 (code string) string{
	Md5 :=md5.New()
	io.WriteString(Md5,code)
	return hex.EncodeToString(Md5.Sum(nil))
}
