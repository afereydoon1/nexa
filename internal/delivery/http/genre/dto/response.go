package dto

type GenreResponse struct {
	ID              uint   `json:"id"`
	Name            string `json:"name"`
	Slug            string `json:"slug"`
	ImageBackground string `json:"image_background"`
}
