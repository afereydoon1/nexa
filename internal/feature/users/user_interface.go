package users

import (
	"nexa/internal/feature/users/domain"
)

type UserRepository interface {
	Create(User *domain.User) error
	FindByEmail(email string) (*domain.User, error)
}
