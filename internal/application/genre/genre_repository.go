package genre

import "nexa/internal/domain"

type GenreRepository interface {

	//Genre List
	GetAll() ([]domain.Genre, error)

	// Create Genre
	Create(genre *domain.Genre) error

	//Get Single Genre
	FindByID(id uint) (*domain.Genre, error)

	//Update Genre
	Update(genre *domain.Genre) error

	//Delete Genre
	Delete(id uint) error
}
