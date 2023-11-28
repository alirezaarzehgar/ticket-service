package database

import (
	"gorm.io/gorm"

	"github.com/alirezaarzehgar/ticketservice/config"
	"github.com/alirezaarzehgar/ticketservice/model"
	"github.com/alirezaarzehgar/ticketservice/util"
)

func usersMigrate(db *gorm.DB) error {
	c := config.Admin()
	admin := model.User{
		Username: c.Username,
		Email:    c.Email,
		Password: util.CreateSHA256(c.Password),
		Role:     model.USERS_ROLE_ADMIN,
	}
	return db.Create(&admin).Error
}

func Migrate(db *gorm.DB) error {
	if db.Migrator().HasTable(&model.User{}) {
		return nil
	}

	err := db.AutoMigrate(&model.User{})
	if err != nil {
		return err
	}

	if err := usersMigrate(db); err != nil {
		return err
	}

	return nil
}
