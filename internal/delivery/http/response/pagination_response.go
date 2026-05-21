package response

import (
	i18n "nexa/internal/infra/lang"

	"github.com/gin-gonic/gin"
)

type PaginationMeta struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
	Total int `json:"total"`
}

func PaginatedResponse(
	c *gin.Context,
	status int,
	code string,
	data interface{},
	meta PaginationMeta,
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
		Meta:    meta,
	})
}
