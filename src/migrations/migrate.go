package migrations

import (
	"github.com/sahidhossen/synmail/src/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	// Migrate multiple tables
	err := db.AutoMigrate(
		&models.User{},
		&models.Campaign{},
		&models.Subscribe{},
		&models.Trackers{},
	)
	if err != nil {
		panic("Migration failed: " + err.Error())
	}
}
