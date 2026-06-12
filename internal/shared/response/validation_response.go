package response

import (
	i18n "nexa/internal/shared/lang"

	"github.com/gin-gonic/gin"
)

func ValidationErrorResponse(
	c *gin.Context,
	errors map[string]string,
	tr *i18n.Translator,
) {

	lang := c.GetHeader("Accept-Language")

	if lang == "" {
		lang = "en"
	}

	result := map[string]string{}

	for field, code := range errors {
		result[field] = tr.Translate(lang, code)
	}

	c.JSON(400, APIResponse{
		Success: false,
		Message: "validation failed",
		Errors:  result,
	})
}
