package users

import (
	appErr "nexa/internal/shared/errors"
)

var RegisterValidationMessages = map[string]map[string]string{

	"name": {
		"required": appErr.ErrNameRequired,
		"min":      appErr.ErrNameMin,
	},

	"email": {
		"required": appErr.ErrEmailRequired,
		"email":    appErr.ErrEmailInvalid,
	},

	"password": {
		"required": appErr.ErrPasswordRequired,
		"min":      appErr.ErrPasswordMin,
	},
}

var LoginValidationMessages = map[string]map[string]string{

	"email": {
		"required": appErr.ErrEmailRequired,
		"email":    appErr.ErrEmailInvalid,
	},

	"password": {
		"required": appErr.ErrPasswordRequired,
	},
}
