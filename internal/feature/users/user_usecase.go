package users

import (
	"errors"
	"nexa/internal/feature/users/domain"
	"nexa/internal/shared/security"
)

type UserUseCase struct {
	repo         UserRepository
	tokenService TokenService
}

func NewUserUseCase(r UserRepository, t TokenService) *UserUseCase {
	return &UserUseCase{
		repo:         r,
		tokenService: t,
	}
}

func (uc *UserUseCase) Create(
	name string,
	email string,
	password string,
) (*domain.User, error) {

	existingUser, _ := uc.repo.FindByEmail(email)
	if existingUser != nil {
		return nil, errors.New("user already exists")
	}

	hashedPassword, err := security.HashPassword(password)
	if err != nil {
		return nil, err
	}

	user := &domain.User{
		Name:     name,
		Email:    email,
		Password: hashedPassword,
	}

	err = uc.repo.Create(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (uc *UserUseCase) Login(
	email string,
	password string,
) (string, *domain.User, error) {

	user, err := uc.repo.FindByEmail(email)
	if err != nil {
		return "", nil, errors.New("invalid credentials")
	}

	err = security.CheckPassword(password, user.Password)
	if err != nil {
		return "", nil, errors.New("invalid credentials")
	}

	token, err := uc.tokenService.GenerateToken(
		user.ID,
		user.Email,
	)

	if err != nil {
		return "", nil, err
	}

	return token, user, nil
}
