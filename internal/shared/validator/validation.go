package validator

import (
	"errors"
	"strings"

	v "github.com/go-playground/validator/v10"
)

func FormatValidationError(
	err error,
	messages map[string]map[string]string,
) map[string]string {

	result := map[string]string{}

	var ve v.ValidationErrors

	if errors.As(err, &ve) {

		for _, e := range ve {

			field := strings.ToLower(e.Field())
			tag := e.Tag()

			fieldRules, ok := messages[field]
			if !ok {
				continue
			}

			code, exists := fieldRules[tag]
			if exists {
				result[field] = code
			}
		}
	}

	return result
}
