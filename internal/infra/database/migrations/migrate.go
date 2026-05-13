package migrations

import (
	"nexa/internal/infra/database/models"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {

	return db.AutoMigrate(
		&models.UserModel{},
		&models.GenreModel{},
	)
}
