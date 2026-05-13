package repository

import (
	"nexa/internal/application/genre"
	"nexa/internal/domain"
	"nexa/internal/infra/database/models"

	"gorm.io/gorm"
)

type genreRepository struct {
	db *gorm.DB
}

func NewGenreRepository(db *gorm.DB) genre.GenreRepository {
	return &genreRepository{
		db: db,
	}
}

func (r *genreRepository) GetAll() ([]domain.Genre, error) {

	var modelsList []models.GenreModel

	err := r.db.Find(&modelsList).Error
	if err != nil {
		return nil, err
	}

	var genres []domain.Genre

	for _, model := range modelsList {

		genres = append(genres, domain.Genre{
			ID:              model.ID,
			Name:            model.Name,
			Slug:            model.Slug,
			ImageBackground: model.ImageBackground,
		})
	}

	return genres, nil
}

func (r *genreRepository) Create(genreData *domain.Genre) error {

	model := models.GenreModel{
		Name:            genreData.Name,
		Slug:            genreData.Slug,
		ImageBackground: genreData.ImageBackground,
	}

	return r.db.Create(&model).Error
}

func (r *genreRepository) FindByID(id uint) (*domain.Genre, error) {

	var model models.GenreModel

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

	model := models.GenreModel{
		ID:              genreData.ID,
		Name:            genreData.Name,
		Slug:            genreData.Slug,
		ImageBackground: genreData.ImageBackground,
	}

	return r.db.Save(&model).Error
}

func (r *genreRepository) Delete(id uint) error {

	return r.db.Delete(&models.GenreModel{}, id).Error
}
