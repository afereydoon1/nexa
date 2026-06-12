package genres

import (
	"nexa/internal/feature/genres/domain"

	"gorm.io/gorm"
)

type genreRepository struct {
	db *gorm.DB
}

func NewGenreRepository(db *gorm.DB) GenreRepository {
	return &genreRepository{
		db: db,
	}
}

func (r *genreRepository) GetAll() ([]domain.Genre, error) {

	var modelsList []GenreModel

	err := r.db.Find(&modelsList).Error
	if err != nil {
		return nil, err
	}

	var genreList []domain.Genre

	for _, model := range modelsList {

		genreList = append(genreList, domain.Genre{
			ID:              model.ID,
			Name:            model.Name,
			Slug:            model.Slug,
			ImageBackground: model.ImageBackground,
		})
	}

	return genreList, nil
}

func (r *genreRepository) Create(genreData *domain.Genre) error {

	model := GenreModel{
		Name:            genreData.Name,
		Slug:            genreData.Slug,
		ImageBackground: genreData.ImageBackground,
	}

	return r.db.Create(&model).Error
}

func (r *genreRepository) FindByID(id uint) (*domain.Genre, error) {

	var model GenreModel

	err := r.db.
		Where("id = ?", id).
		First(&model).Error

	if err != nil {
		return nil, err
	}

	return &domain.Genre{
		ID:              model.ID,
		Name:            model.Name,
		Slug:            model.Slug,
		ImageBackground: model.ImageBackground,
	}, nil
}

func (r *genreRepository) Update(genreData *domain.Genre) error {

	model := GenreModel{
		ID:              genreData.ID,
		Name:            genreData.Name,
		Slug:            genreData.Slug,
		ImageBackground: genreData.ImageBackground,
	}

	return r.db.Save(&model).Error
}

func (r *genreRepository) Delete(id uint) error {

	return r.db.Delete(&GenreModel{}, id).Error
}
