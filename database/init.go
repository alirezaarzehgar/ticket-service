package database

import (
	"fmt"
	"log/slog"

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

	slog.Debug("database dsn", "data", dsn)

	newLogger := logger.New(writer, logger.Config{LogLevel: logger.Info})
	return gorm.Open(mysql.Open(dsn), &gorm.Config{TranslateError: true, Logger: newLogger})
}
