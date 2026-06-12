package users

import (
	"net/http"
	"nexa/internal/feature/users/dto"
	appErr "nexa/internal/shared/errors"
	i18n "nexa/internal/shared/lang"
	"nexa/internal/shared/response"
	"nexa/internal/shared/validator"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	usecase    *UserUseCase
	translator *i18n.Translator
}

func NewUserHandler(uc *UserUseCase, translator *i18n.Translator) *UserHandler {
	return &UserHandler{
		usecase:    uc,
		translator: translator,
	}
}

func (h *UserHandler) Register(c *gin.Context) {

	var req dto.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		validationErrors := validator.FormatValidationError(
			err,
			RegisterValidationMessages,
		)

		response.ValidationErrorResponse(
			c,
			validationErrors,
			h.translator,
		)

		return
	}

	userData, err := h.usecase.Create(
		req.Name,
		req.Email,
		req.Password,
	)

	if err != nil {

		response.ErrorResponse(
			c,
			http.StatusBadRequest,
			appErr.ErrUserCreateFailed,
			h.translator,
		)

		return
	}

	response.SuccessResponse(
		c,
		http.StatusCreated,
		appErr.SuccessUserCreated,
		userData,
		h.translator,
	)
}

func (h *UserHandler) Login(c *gin.Context) {

	var req dto.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		validationErrors := validator.FormatValidationError(
			err,
			LoginValidationMessages,
		)

		response.ValidationErrorResponse(
			c,
			validationErrors,
			h.translator,
		)

		return
	}

	token, userData, err := h.usecase.Login(
		req.Email,
		req.Password,
	)

	if err != nil {

		response.ErrorResponse(
			c,
			http.StatusUnauthorized,
			appErr.ErrInvalidCredentials,
			h.translator,
		)

		return
	}

	response.SuccessResponse(
		c,
		http.StatusOK,
		appErr.SuccessLogin,
		dto.LoginResponse{
			Token: token,
			User: dto.UserResponse{
				ID:    userData.ID,
				Name:  userData.Name,
				Email: userData.Email,
			},
		},
		h.translator,
	)
}
