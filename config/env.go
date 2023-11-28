package config

import (
	"os"
	"strconv"

	"gorm.io/gorm/logger"
)

func ListenerAddr() string {
	return os.Getenv("RUNNING_ADDR")
}

func JwtSecret() []byte {
	return []byte(os.Getenv("JWT_SECRET"))
}

type DbConf struct {
	Host     string
	User     string
	Password string
	DbName   string
	Port     uint64
	DbLogger logger.Writer
}

func GetDb() (*DbConf, error) {
	port, err := strconv.ParseUint(os.Getenv("MYSQL_PORT"), 10, 64)
	if err != nil {
		return nil, err
	}

	conf := DbConf{
		Host:     os.Getenv("MYSQL_HOST"),
		User:     os.Getenv("MYSQL_USER"),
		Password: os.Getenv("MYSQL_PASSWORD"),
		DbName:   os.Getenv("MYSQL_DATABASE"),
		Port:     port,
		DbLogger: nil,
	}
	return &conf, nil
}

func AlertDb() string {
	return os.Getenv("ALERT_DATABASE")
}

func Debug() bool {
	v, err := strconv.ParseBool(os.Getenv("DEBUG"))
	if err != nil {
		return false
	}
	return v
}

type AdminConfig struct {
	Username, Email, Password string
}

func Admin() AdminConfig {
	return AdminConfig{
		Username: os.Getenv("ADMIN_NAME"),
		Email:    os.Getenv("ADMIN_EMAIL"),
		Password: os.Getenv("ADMIN_PASSWORD"),
	}
}
