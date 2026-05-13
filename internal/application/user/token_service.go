package user

type TokenService interface {
	GenerateToken(userID uint, email string) (string, error)
}
