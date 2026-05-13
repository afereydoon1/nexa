package user

import (
	"nexa/internal/domain"
)

type UserRepository interface {
	Create(User *domain.User) error
	FindByEmail(email string) (*domain.User, error)
}
