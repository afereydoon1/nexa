package http

import (
	"nexa/internal/delivery/http/user/dto"
	"net/http"

	"nexa/internal/application/user"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	usecase *user.UserUseCase
}

func NewUserHandler(uc *user.UserUseCase) *UserHandler {
	return &UserHandler{
		usecase: uc,
	}
}

func (h *UserHandler) Register(c *gin.Context) {

	var req dto.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := h.usecase.Create(
		req.Name,
		req.Email,
		req.Password,
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "user created",
	})
}

func (h *UserHandler) Login(c *gin.Context) {

	var req dto.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	token, err := h.usecase.Login(
		req.Email,
		req.Password,
	)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.LoginResponse{
		Token: token,
	})
}
