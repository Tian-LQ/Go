package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var GlobalDB *gorm.DB

func init() {
	db, _ := gorm.Open(mysql.New(mysql.Config{
		DSN: "root:tlq19971019@tcp(127.0.0.1:3306)/gorm_database?charset=utf8mb4&parseTime=True&loc=Local",
	}), &gorm.Config{
		SkipDefaultTransaction:                   false,
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	// 连接池相关配置
	sqlDB, _ := db.DB()
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	// 一切准备就绪，初始化全局变量GlobalDB
	GlobalDB = db
}
