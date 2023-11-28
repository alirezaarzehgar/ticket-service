package database

import (
	"gorm.io/gorm"

	"github.com/alirezaarzehgar/ticketservice/model"
)

func Migrate(db *gorm.DB, defaultAdmin model.User) error {
	if db.Migrator().HasTable(&model.User{}) {
		return nil
	}

	err := db.AutoMigrate(&model.User{})
	if err != nil {
		return err
	}

	if err := db.Create(&defaultAdmin).Error; err != nil {
		return err
	}

	return nil
}
