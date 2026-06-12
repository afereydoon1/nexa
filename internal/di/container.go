package di

import (
	"nexa/internal/feature/genres"
	"nexa/internal/feature/users"
	"nexa/internal/shared/auth"
	"nexa/internal/shared/config"
	i18n "nexa/internal/shared/lang"
	localStorage "nexa/pkg/storage/local"

	"gorm.io/gorm"
)

type AppHandlers struct {
	UserHandler  *users.UserHandler
	GenreHandler *genres.GenreHandler
}

func InitHandlers(db *gorm.DB, cfg *config.Config) *AppHandlers {

	// Services
	jwtService := auth.NewJWTService(cfg.JWTSecret)
	storageService := localStorage.NewStorageService()
	translator := i18n.NewTranslator()

	// User
	userRepo := users.NewUserRepository(db)
	userUseCase := users.NewUserUseCase(userRepo, jwtService)
	userHandler := users.NewUserHandler(userUseCase, translator)

	//Genres
	genreRepo := genres.NewGenreRepository(db)
	genreUseCase := genres.NewGenreUseCase(genreRepo)
	genreHandler := genres.NewGenreHandler(genreUseCase, storageService, translator)

	return &AppHandlers{
		UserHandler:  userHandler,
		GenreHandler: genreHandler,
	}
}
