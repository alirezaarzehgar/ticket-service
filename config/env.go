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
		Role:     model.USERS_ROLE_ADMIN,
	}
}
