package config

import (
	"crypto/sha256"
	"encoding/hex"
	"os"
	"strconv"

	"github.com/alirezaarzehgar/ticketservice/database"
	"github.com/alirezaarzehgar/ticketservice/model"
)

func ListenerAddr() string {
	return os.Getenv("RUNNING_ADDR")
}

func JwtSecret() []byte {
	return []byte(os.Getenv("JWT_SECRET"))
}

func GetDb() (*database.DbConf, error) {
	port, err := strconv.ParseUint(os.Getenv("MYSQL_PORT"), 10, 64)
	if err != nil {
		return nil, err
	}

	conf := database.DbConf{
		Host:     os.Getenv("MYSQL_HOST"),
		User:     os.Getenv("MYSQL_USER"),
		Password: os.Getenv("MYSQL_PASSWORD"),
		DbName:   os.Getenv("MYSQL_DATABASE"),
		Port:     port,
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

func Admin() model.User {
	hashByte := sha256.Sum256([]byte(os.Getenv("ADMIN_PASSWORD")))
	passStr := hex.EncodeToString(hashByte[:])

	return model.User{
		Username: os.Getenv("ADMIN_NAME"),
		Email:    os.Getenv("ADMIN_EMAIL"),
		Password: passStr,
		Role:     model.USERS_ROLE_SUPER_ADMIN,
	}
}

func Assets() string {
	dir := os.Getenv("ASSETS_DIRECTORY")
	if _, err := os.Stat(dir); err != nil {
		os.MkdirAll(dir, os.ModePerm)
	}
	return dir
}

type SmtpConf struct {
	FromAddress string
	FromName    string
	Password    string
	Host        string
	Port        string
	Server      string
}

func GetMailConfig() SmtpConf {
	return SmtpConf{
		FromAddress: os.Getenv("MAIL_FROM_ADDRESS"),
		FromName:    os.Getenv("MAIL_FROM_NAME"),
		Password:    os.Getenv("MAIL_PASSWORD"),
		Host:        os.Getenv("MAIL_HOST"),
		Port:        os.Getenv("MAIL_PORT"),
		Server:      os.Getenv("MAIL_SERVER"),
	}
}
