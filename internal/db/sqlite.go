package db

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func ConnectDB() {
	if DB != nil {
		return
	}
	var err error
	DB, err = gorm.Open(sqlite.Open("larvauth.db"), &gorm.Config{})
	if err != nil {
		log.Panic("连接至数据库失败：", err)
	}
}

func GetDB() *gorm.DB {
	ConnectDB()
	if DB == nil {
		panic("初始化了吗？")
	}
	return DB
}
