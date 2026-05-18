package validator

import (
	"errors"
	"strings"

	appErr "nexa/internal/domain/errors"

	v "github.com/go-playground/validator/v10"
)

func FormatValidationError(err error) map[string]string {

	result := map[string]string{}

	var ve v.ValidationErrors

	if errors.As(err, &ve) {

		for _, e := range ve {

			field := strings.ToLower(e.Field())

			switch field {

			case "name":

				switch e.Tag() {

				case "required":
					result[field] = appErr.ErrNameRequired

				case "min":
					result[field] = appErr.ErrNameMin
				}
			}
		}
	}

	return result
}
