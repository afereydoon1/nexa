package genres

import (
	"net/http"
	"nexa/internal/feature/genres/dto"
	appErr "nexa/internal/shared/errors"
	i18n "nexa/internal/shared/lang"
	"nexa/internal/shared/response"
	"nexa/internal/shared/validator"
	localStorage "nexa/pkg/storage/local"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GenreHandler struct {
	usecase    *GenreUseCase
	storage    *localStorage.StorageService
	translator *i18n.Translator
}

func NewGenreHandler(
	uc *GenreUseCase,
	storage *localStorage.StorageService,
	translator *i18n.Translator,
) *GenreHandler {

	return &GenreHandler{
		usecase:    uc,
		storage:    storage,
		translator: translator,
	}
}

func (h *GenreHandler) Create(c *gin.Context) {

	var req dto.CreateGenreRequest

	if err := c.ShouldBind(&req); err != nil {

		validationErrors := validator.FormatValidationError(
			err,
			dto.CreateGenreValidationMessages,
		)

		response.ValidationErrorResponse(
			c,
			validationErrors,
			h.translator,
		)

		return
	}

	file, err := c.FormFile("image")
	if err != nil {

		response.ErrorResponse(
			c,
			http.StatusBadRequest,
			appErr.ErrImageRequired,
			h.translator,
		)

		return
	}

	imagePath, err := h.storage.Save(
		c,
		file,
		"genres",
	)

	if err != nil {

		response.ErrorResponse(
			c,
			http.StatusInternalServerError,
			appErr.ErrUploadFailed,
			h.translator,
		)

		return
	}

	genre, err := h.usecase.Create(
		req.Name,
		req.Slug,
		imagePath,
	)

	if err != nil {

		response.ErrorResponse(
			c,
			http.StatusBadRequest,
			appErr.ErrGenreCreateFailed,
			h.translator,
		)

		return
	}

	response.SuccessResponse(
		c,
		http.StatusCreated,
		appErr.SuccessGenreCreated,
		genre, // 👈 مهم: دیتا برگشتی
		h.translator,
	)
}

func (h *GenreHandler) GetAll(c *gin.Context) {

	genres, err := h.usecase.GetAll()
	if err != nil {

		response.ErrorResponse(
			c,
			http.StatusInternalServerError,
			appErr.ErrGenreNotFound,
			h.translator,
		)

		return
	}

	result := make([]dto.GenreResponse, 0, len(genres))

	for _, genre := range genres {
		result = append(result, dto.GenreResponse{
			ID:              genre.ID,
			Name:            genre.Name,
			Slug:            genre.Slug,
			ImageBackground: genre.ImageBackground,
		})
	}

	response.SuccessResponse(
		c,
		http.StatusOK,
		appErr.SuccessGenresFetched,
		result,
		h.translator,
	)
}

func (h *GenreHandler) GetByID(c *gin.Context) {

	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {

		response.ErrorResponse(
			c,
			http.StatusBadRequest,
			appErr.ErrInvalidID,
			h.translator,
		)

		return
	}

	genreData, err := h.usecase.FindByID(uint(id))
	if err != nil {

		response.ErrorResponse(
			c,
			http.StatusNotFound,
			appErr.ErrGenreNotFound,
			h.translator,
		)

		return
	}

	result := dto.GenreResponse{
		ID:              genreData.ID,
		Name:            genreData.Name,
		Slug:            genreData.Slug,
		ImageBackground: genreData.ImageBackground,
	}

	response.SuccessResponse(
		c,
		http.StatusOK,
		appErr.SuccessGenreFetched,
		result,
		h.translator,
	)
}

func (h *GenreHandler) Update(c *gin.Context) {

	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {

		response.ErrorResponse(
			c,
			http.StatusBadRequest,
			appErr.ErrInvalidID,
			h.translator,
		)

		return
	}

	var req dto.UpdateGenreRequest

	if err := c.ShouldBind(&req); err != nil {

		validationErrors := validator.FormatValidationError(
			err,
			dto.UpdateGenreValidationMessages,
		)

		response.ValidationErrorResponse(
			c,
			validationErrors,
			h.translator,
		)

		return
	}

	oldGenre, err := h.usecase.FindByID(uint(id))
	if err != nil {

		response.ErrorResponse(
			c,
			http.StatusNotFound,
			appErr.ErrGenreNotFound,
			h.translator,
		)

		return
	}

	imagePath := oldGenre.ImageBackground

	file, err := c.FormFile("image")

	if err == nil {

		imagePath, err = h.storage.Save(
			c,
			file,
			"genres",
		)

		if err != nil {

			response.ErrorResponse(
				c,
				http.StatusInternalServerError,
				appErr.ErrUploadFailed,
				h.translator,
			)

			return
		}

		_ = h.storage.Delete(oldGenre.ImageBackground)
	}

	updatedGenre, err := h.usecase.Update(
		uint(id),
		req.Name,
		req.Slug,
		imagePath,
	)

	if err != nil {

		response.ErrorResponse(
			c,
			http.StatusBadRequest,
			appErr.ErrGenreCreateFailed,
			h.translator,
		)

		return
	}

	response.SuccessResponse(
		c,
		http.StatusOK,
		appErr.SuccessGenreUpdated,
		updatedGenre,
		h.translator,
	)
}

func (h *GenreHandler) Delete(c *gin.Context) {

	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {

		response.ErrorResponse(
			c,
			http.StatusBadRequest,
			appErr.ErrInvalidID,
			h.translator,
		)

		return
	}

	err = h.usecase.Delete(uint(id))
	if err != nil {

		response.ErrorResponse(
			c,
			http.StatusBadRequest,
			appErr.ErrGenreNotFound,
			h.translator,
		)

		return
	}

	response.SuccessResponse(
		c,
		http.StatusOK,
		appErr.SuccessGenreDeleted,
		nil,
		h.translator,
	)
}
