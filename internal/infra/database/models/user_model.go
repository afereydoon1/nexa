package models

type UserModel struct {
	ID       uint `gorm:"primaryKey"`
	Name     string
	Email    string
	Password string
}
