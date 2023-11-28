package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbConf struct {
	Host     string
	User     string
	Password string
	DbName   string
	Port     uint64
}

func Init(c *DbConf, writer logger.Writer) (db *gorm.DB, err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.User, c.Password, c.Host, c.Port, c.DbName)

	newLogger := logger.New(writer, logger.Config{LogLevel: logger.Warn})
	return gorm.Open(mysql.Open(dsn), &gorm.Config{TranslateError: true, Logger: newLogger})
}
