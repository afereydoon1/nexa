package response

import (
	i18n "nexa/internal/shared/lang"

	"github.com/gin-gonic/gin"
)

func ErrorResponse(
	c *gin.Context,
	status int,
	code string,
	tr *i18n.Translator,
) {

	lang := c.GetHeader("Accept-Language")

	if lang == "" {
		lang = "en"
	}

	c.JSON(status, APIResponse{
		Success: false,
		Message: tr.Translate(lang, code),
	})
}
