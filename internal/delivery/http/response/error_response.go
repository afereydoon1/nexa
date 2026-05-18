package response

import (
	i18n "nexa/internal/infra/lang"

	"github.com/gin-gonic/gin"
)

func ValidationErrorResponse(
	c *gin.Context,
	errs map[string]string,
	tr *i18n.Translator,
) {

	lang := c.GetHeader("Accept-Language")

	if lang == "" {
		lang = "en"
	}

	result := map[string]string{}

	for field, code := range errs {
		result[field] = tr.Translate(lang, code)
	}

	c.JSON(400, gin.H{
		"errors": result,
	})
}
