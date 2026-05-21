package http

import (
	"nexa/internal/delivery/http/response"
	userValidation "nexa/internal/delivery/http/user"
	"nexa/internal/delivery/http/user/dto"
	"nexa/internal/delivery/http/validator"
	appErr "nexa/internal/domain/errors"
	"net/http"

	"nexa/internal/application/user"
	i18n "nexa/internal/infra/lang"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	usecase    *user.UserUseCase
	translator *i18n.Translator
}

func NewUserHandler(uc *user.UserUseCase, translator *i18n.Translator) *UserHandler {
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
			userValidation.RegisterValidationMessages,
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
			userValidation.LoginValidationMessages,
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
