package main

import (
	"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

var db *gorm.DB

func initDb() {
	db, err := gorm.Open("mysql", "root:123456@(localhost)/TEST?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
		return
	}
	//开启日志模式
	db.LogMode(true)
	// 获取通用数据库对象 sql.DB ，然后使用其提供的功能
	sqlDB := db.DB()
	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

}
