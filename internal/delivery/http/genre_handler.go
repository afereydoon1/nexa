package http

import (
	application "nexa/internal/application/genre"
	"nexa/internal/delivery/http/genre/dto"
	localStorage "nexa/pkg/storage/local"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GenreHandler struct {
	usecase *application.GenreUseCase
	storage *localStorage.StorageService
}

func NewGenreHandler(
	uc *application.GenreUseCase,
	storage *localStorage.StorageService,
) *GenreHandler {

	return &GenreHandler{
		usecase: uc,
		storage: storage,
	}
}

func (h *GenreHandler) Create(c *gin.Context) {

	var req dto.CreateGenreRequest

	if err := c.ShouldBind(&req); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	file, err := c.FormFile("image")
	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "image is required",
		})
		return
	}

	imagePath, err := h.storage.Save(
		c,
		file,
		"genres",
	)

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

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
			"error": "invalid id",
		})
		return
	}

	var req dto.UpdateGenreRequest

	if err := c.ShouldBind(&req); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// -----------------------
	// get old genre (important for image delete)
	// -----------------------
	oldGenre, err := h.usecase.FindByID(uint(id))
	if err != nil {

		c.JSON(http.StatusNotFound, gin.H{
			"error": "genre not found",
		})
		return
	}

	// handle image upload
	file, err := c.FormFile("image")

	imagePath := oldGenre.ImageBackground // default = keep old image

	if err == nil {
		// upload new images
		imagePath, err = h.storage.Save(
			c,
			file,
			"genres",
		)

		if err != nil {

			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		// Delete old image
		_ = h.storage.Delete(oldGenre.ImageBackground)
	}

	// update usecase
	err = h.usecase.Update(
		uint(id),
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
