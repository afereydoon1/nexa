package models

type GenreModel struct {
	ID              uint `gorm:"primaryKey"`
	Name            string
	Slug            string
	ImageBackground string
}
