package database

import (
	"Music-library/internal/models"
	myLogger "Music-library/pkg/logger"
)

func Migrate(logger *myLogger.Logger) {
	err := DB.AutoMigrate(&models.Song{})
	if err != nil {
		logger.Err.Fatalf("Error migrating database: %v", err)
	}
	logger.Info.Println("Migrated database successfully")
}
