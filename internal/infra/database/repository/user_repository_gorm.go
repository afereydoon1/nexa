package repository

import (
	"nexa/internal/application/user"
	"nexa/internal/domain"
	"nexa/internal/infra/database/models"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) user.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) Create(user *domain.User) error {

	model := models.UserModel{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}

	return r.db.Create(&model).Error
}

func (r *userRepository) FindByEmail(email string) (*domain.User, error) {

	var model models.UserModel

	err := r.db.
		Where("email = ?", email).
		First(&model).Error

	if err != nil {
		return nil, err
	}

	return &domain.User{
		ID:       model.ID,
		Name:     model.Name,
		Email:    model.Email,
		Password: model.Password,
	}, nil
}
