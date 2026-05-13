package router

import (
	"nexa/internal/di"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, handlers *di.AppHandlers) {

	// API Prefix
	api := r.Group("/api/v1")

	//auth
	users := api.Group("/auth")
	{
		users.POST("/register", handlers.UserHandler.Register)
		users.POST("/login", handlers.UserHandler.Login)
	}

	//genres
	genres := api.Group("/genres")
	{
		genres.GET("/", handlers.GenreHandler.GetAll)
		genres.POST("/", handlers.GenreHandler.Create)
		genres.GET("/:id", handlers.GenreHandler.GetByID)
		genres.PUT("/:id", handlers.GenreHandler.Update)
		genres.DELETE("/:id", handlers.GenreHandler.Delete)
	}
}
