package middleware

import (
	"gorm.io/gorm"
)

var db *gorm.DB

func SetDB(externalDb *gorm.DB) {
	db = externalDb
}
