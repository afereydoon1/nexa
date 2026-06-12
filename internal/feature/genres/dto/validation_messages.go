package dto

import (
	appErr "nexa/internal/shared/errors"
)

var CreateGenreValidationMessages = map[string]map[string]string{

	"name": {
		"required": appErr.ErrNameRequired,
		"min":      appErr.ErrNameMin,
	},

	"slug": {
		"required": appErr.ErrSlugRequired,
	},
}

var UpdateGenreValidationMessages = map[string]map[string]string{

	"name": {
		"required": appErr.ErrNameRequired,
		"min":      appErr.ErrNameMin,
	},

	"slug": {
		"required": appErr.ErrSlugRequired,
	},
}
