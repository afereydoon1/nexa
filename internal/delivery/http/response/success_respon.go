package response

import (
	i18n "nexa/internal/infra/lang"

	"github.com/gin-gonic/gin"
)

func SuccessResponse(
	c *gin.Context,
	status int,
	code string,
	data interface{},
	tr *i18n.Translator,
) {

	lang := c.GetHeader("Accept-Language")

	if lang == "" {
		lang = "en"
	}

	c.JSON(status, APIResponse{
		Success: true,
		Message: tr.Translate(lang, code),
		Data:    data,
	})
}
