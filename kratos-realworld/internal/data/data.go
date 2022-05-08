package data

import (
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"kratos-realworld/internal/conf"
	"time"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewDB, NewUserRepo, NewProfileRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	db *gorm.DB
}

// NewData .
func NewData(c *conf.Data, logger log.Logger, db *gorm.DB) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{db: db}, cleanup, nil
}

func NewDB(c *conf.Data) *gorm.DB {
	if c != nil {
		fmt.Printf("the mysql dsn: %s\n", c.Database.Dsn)
	}
	db, _ := gorm.Open(mysql.New(mysql.Config{
		DSN: "root:tlq19971019@tcp(127.0.0.1:3306)/rwdb?charset=utf8mb4&parseTime=True&loc=Local",
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
	db.AutoMigrate(&User{})
	return db
}
