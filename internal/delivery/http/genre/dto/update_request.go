package dto

type UpdateGenreRequest struct {
	Name string `form:"name" binding:"required,min=3"`
	Slug string `form:"slug" binding:"required"`
}
