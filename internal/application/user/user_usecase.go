package user

import (
	"errors"

	"nexa/internal/domain"
	"nexa/internal/infra/security"
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
) error {

	// Check existing user
	existingUser, _ := uc.repo.FindByEmail(email)

	if existingUser != nil {
		return errors.New("email already exists")
	}

	// Hash password
	hashedPassword, err := security.HashPassword(password)
	if err != nil {
		return err
	}

	user := &domain.User{
		Name:     name,
		Email:    email,
		Password: hashedPassword,
	}

	return uc.repo.Create(user)
}

func (uc *UserUseCase) Login(
	email string,
	password string,
) (string, error) {

	user, err := uc.repo.FindByEmail(email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	err = security.CheckPassword(
		password,
		user.Password,
	)

	if err != nil {
		return "", errors.New("invalid credentials")
	}

	token, err := uc.tokenService.GenerateToken(
		user.ID,
		user.Email,
	)

	if err != nil {
		return "", err
	}

	return token, nil
}
