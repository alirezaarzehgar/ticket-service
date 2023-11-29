package database

import (
	"log/slog"

	"gorm.io/gorm"

	"github.com/alirezaarzehgar/ticketservice/model"
)

func Migrate(db *gorm.DB, defaultAdmin model.User) error {
	if db.Migrator().HasTable(&model.User{}) {
		slog.Debug("skip migration")
		return nil
	}

	err := db.AutoMigrate(&model.User{})
	if err != nil {
		return err
	}
	slog.Debug("migrate tables")

	if err := db.Create(&defaultAdmin).Error; err != nil {
		return err
	}
	slog.Debug("create default admin", "data", defaultAdmin)

	return nil
}
