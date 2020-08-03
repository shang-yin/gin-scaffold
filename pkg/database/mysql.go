package database

import (
	"fmt"
	"gin-scaffold/pkg/conf"
	"strconv"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func init() {
	var err error
	if ok := conf.Config.IsSet("database"); !ok {
		panic("Please set mysql link parameters first")
	}
	database := conf.Config.GetStringMapString("database")
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=%v&parseTime=True&loc=Local", database["username"], database["password"],
		database["server"], database["port"], database["dbname"], database["charset"])

	DB, err = gorm.Open(mysql.New(mysql.Config{
		DriverName: database["driver"],
		DSN:        dsn,
	}), &gorm.Config{NamingStrategy: schema.NamingStrategy{
		TablePrefix:   database["prefix"],
		SingularTable: true,
	}})
	if err != nil {
		panic("connect mysql failed")
	}
	sqlDB, err := DB.DB()
	if err != nil {
		panic("connect mysql failed")
	}
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	maxIdleConn, _ := strconv.Atoi(database["maxIdleConn"])
	sqlDB.SetMaxIdleConns(maxIdleConn)

	// SetMaxOpenConns 设置打开数据库连接的最大数量
	maxOpenConn, _ := strconv.Atoi(database["maxOpenConn"])
	sqlDB.SetMaxOpenConns(maxOpenConn)

	// SetConnMaxLifetime 设置了连接可复用的最大时间
	maxLifeTime, _ := strconv.Atoi(database["maxLifeTime"])
	duration, _ := time.ParseDuration(strconv.Itoa(maxLifeTime))
	sqlDB.SetConnMaxLifetime(duration * time.Second)

	if err := sqlDB.Ping(); err != nil {
		panic(err)
	}
}
