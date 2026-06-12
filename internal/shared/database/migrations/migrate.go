package migrations

import (
	"nexa/internal/feature/genres"
	"nexa/internal/feature/users"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {

	return db.AutoMigrate(
		&users.UserModel{},
		&genres.GenreModel{},
	)
}
