package genre

import "nexa/internal/domain"

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
) error {

	genre := &domain.Genre{
		Name:            name,
		Slug:            slug,
		ImageBackground: imageBackground,
	}

	return uc.repo.Create(genre)
}

func (uc *GenreUseCase) Update(
	id uint,
	name string,
	slug string,
	imageBackground string,
) error {

	genre, err := uc.repo.FindByID(id)
	if err != nil {
		return err
	}

	genre.Name = name
	genre.Slug = slug
	genre.ImageBackground = imageBackground

	return uc.repo.Update(genre)
}

func (uc *GenreUseCase) Delete(id uint) error {
	return uc.repo.Delete(id)
}
