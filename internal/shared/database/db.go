package database

import (
	"fmt"
	"nexa/internal/shared/config"
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Connect establishes a database connection based on the provided configuration
// Returns a *gorm.DB instance or an error if connection fails
func Connect(cfg *config.Config) (*gorm.DB, error) {
	dbType := strings.ToLower(cfg.DBConnection)

	var dialector gorm.Dialector
	var dsn string

	switch dbType {
	case "postgres", "pgsql":
		// Build DSN string using SSL mode from config
		dsn = fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
			cfg.DBHost,
			cfg.DBUser,
			cfg.DBPassword,
			cfg.DBName,
			cfg.DBPort,
			cfg.DBSSLMode,
		)
		dialector = postgres.Open(dsn)

	default:
		// Unsupported database type
		return nil, fmt.Errorf("unsupported database type: %s", dbType)
	}

	// Open the Gorm DB connection
	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	// Ping the database to ensure connection is alive
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get sql.DB instance: %w", err)
	}

	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return db, nil
}
