package genres

import (
	"nexa/internal/feature/genres/domain"
)

type GenreUseCase struct {
	repo GenreRepository
}

func NewGenreUseCase(r GenreRepository) *GenreUseCase {
	return &GenreUseCase{
		repo: r,
	}
}

func (uc *GenreUseCase) GetAll() ([]domain.Genre, error) {
	return uc.repo.GetAll()
}

func (uc *GenreUseCase) FindByID(id uint) (*domain.Genre, error) {
	return uc.repo.FindByID(id)
}

func (uc *GenreUseCase) Create(
	name string,
	slug string,
	imageBackground string,
) (*domain.Genre, error) {

	genre := &domain.Genre{
		Name:            name,
		Slug:            slug,
		ImageBackground: imageBackground,
	}

	err := uc.repo.Create(genre)
	if err != nil {
		return nil, err
	}

	return genre, nil
}

func (uc *GenreUseCase) Update(
	id uint,
	name string,
	slug string,
	imageBackground string,
) (*domain.Genre, error) {

	genre, err := uc.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	genre.Name = name
	genre.Slug = slug
	genre.ImageBackground = imageBackground

	err = uc.repo.Update(genre)
	if err != nil {
		return nil, err
	}

	return genre, nil
}

func (uc *GenreUseCase) Delete(id uint) error {
	return uc.repo.Delete(id)
}
