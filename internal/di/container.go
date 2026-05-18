package di

import (
	genreApp "nexa/internal/application/genre"
	userApp "nexa/internal/application/user"
	i18n "nexa/internal/infra/lang"
	localStorage "nexa/pkg/storage/local"

	httpDelivery "nexa/internal/delivery/http"
	"nexa/internal/infra/auth"
	"nexa/internal/infra/config"
	"nexa/internal/infra/database/repository"

	"gorm.io/gorm"
)

type AppHandlers struct {
	UserHandler  *httpDelivery.UserHandler
	GenreHandler *httpDelivery.GenreHandler
}

func InitHandlers(db *gorm.DB, cfg *config.Config) *AppHandlers {

	// Services
	jwtService := auth.NewJWTService(cfg.JWTSecret)
	storageService := localStorage.NewStorageService()
	translator := i18n.NewTranslator()

	// User
	userRepo := repository.NewUserRepository(db)
	userUseCase := userApp.NewUserUseCase(userRepo, jwtService)
	userHandler := httpDelivery.NewUserHandler(userUseCase)

	//Genres
	genreRepo := repository.NewGenreRepository(db)
	genreUseCase := genreApp.NewGenreUseCase(genreRepo)
	genreHandler := httpDelivery.NewGenreHandler(genreUseCase, storageService, translator)

	return &AppHandlers{
		UserHandler:  userHandler,
		GenreHandler: genreHandler,
	}
}
