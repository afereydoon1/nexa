package http

import (
	application "nexa/internal/application/genre"
	"nexa/internal/delivery/http/genre/dto"
	"nexa/internal/delivery/http/response"
	"nexa/internal/delivery/http/validator"
	i18n "nexa/internal/infra/lang"
	localStorage "nexa/pkg/storage/local"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GenreHandler struct {
	usecase *application.GenreUseCase
	storage *localStorage.StorageService
	translator *i18n.Translator
}

func NewGenreHandler(
	uc *application.GenreUseCase,
	storage *localStorage.StorageService,
	translator *i18n.Translator,
) *GenreHandler {

	return &GenreHandler{
		usecase: uc,
		storage: storage,
		translator: translator,
	}
}

func (h *GenreHandler) Create(c *gin.Context) {

	var req dto.CreateGenreRequest

	// validation errors
	if err := c.ShouldBind(&req); err != nil {

		validationErrors := validator.FormatValidationError(err)

		response.ValidationErrorResponse(
			c,
			validationErrors,
			h.translator,
		)

		return
	}

	// image validation
	file, err := c.FormFile("image")
	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"errors": gin.H{
				"image": "image is required",
			},
		})

		return
	}

	// upload image
	imagePath, err := h.storage.Save(
		c,
		file,
		"genres",
	)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to upload image",
		})

		return
	}

	// business logic
	err = h.usecase.Create(
		req.Name,
		req.Slug,
		imagePath,
	)

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "genre created",
	})
}

func (h *GenreHandler) GetAll(c *gin.Context) {

	genres, err := h.usecase.GetAll()

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	var response []dto.GenreResponse

	for _, genre := range genres {

		response = append(response, dto.GenreResponse{
			ID:              genre.ID,
			Name:            genre.Name,
			Slug:            genre.Slug,
			ImageBackground: genre.ImageBackground,
		})
	}

	c.JSON(http.StatusOK, response)
}

func (h *GenreHandler) GetByID(c *gin.Context) {

	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})

		return
	}

	genreData, err := h.usecase.FindByID(uint(id))
	if err != nil {

		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})

		return
	}

	response := dto.GenreResponse{
		ID:              genreData.ID,
		Name:            genreData.Name,
		Slug:            genreData.Slug,
		ImageBackground: genreData.ImageBackground,
	}

	c.JSON(http.StatusOK, response)
}

func (h *GenreHandler) Update(c *gin.Context) {

	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"errors": gin.H{
				"id": "invalid id",
			},
		})
		return
	}

	var req dto.UpdateGenreRequest

	// validation (same system as Create)
	if err := c.ShouldBind(&req); err != nil {

		validationErrors := validator.FormatValidationError(err)

		response.ValidationErrorResponse(
			c,
			validationErrors,
			h.translator,
		)

		return
	}

	// get old genre
	oldGenre, err := h.usecase.FindByID(uint(id))
	if err != nil {

		c.JSON(http.StatusNotFound, gin.H{
			"errors": gin.H{
				"id": "genre not found",
			},
		})
		return
	}

	// handle image upload (same pattern but safer)
	file, err := c.FormFile("image")

	imagePath := oldGenre.ImageBackground

	if err == nil {

		imagePath, err = h.storage.Save(
			c,
			file,
			"genres",
		)

		if err != nil {

			c.JSON(http.StatusBadRequest, gin.H{
				"errors": gin.H{
					"image": "failed to upload image",
				},
			})
			return
		}

		_ = h.storage.Delete(oldGenre.ImageBackground)
	}

	// business logic
	err = h.usecase.Update(
		uint(id),
		req.Name,
		req.Slug,
		imagePath,
	)

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"errors": gin.H{
				"global": err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "genre updated",
	})
}

func (h *GenreHandler) Delete(c *gin.Context) {

	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})

		return
	}

	err = h.usecase.Delete(uint(id))

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "genre deleted",
	})
}
