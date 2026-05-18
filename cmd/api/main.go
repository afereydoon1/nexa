package main

import (
	"nexa/internal/delivery/http/middleware"
	"log"

	"nexa/internal/di"
	"nexa/internal/infra/config"
	dbinfra "nexa/internal/infra/database"
	"nexa/internal/infra/database/migrations"
	"nexa/internal/router"

	"github.com/gin-gonic/gin"
)

func main() {

	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	// Connect database
	db, err := dbinfra.Connect(cfg)
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	log.Println("Database connected successfully")

	// Auto migrate
	if err := migrations.Migrate(db); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	log.Println("Database migrated successfully")

	// Dependency Injection Container
	handlers := di.InitHandlers(db, cfg)

	// Router
	r := gin.Default()

	//Global Middleware
	middleware.SetupCors(r)

	r.Static(
		"/uploads",
		"./storage/uploads",
	)

	// Register Routes
	router.RegisterRoutes(r, handlers)

	// Start server
	log.Println("Server running on :8080")

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
