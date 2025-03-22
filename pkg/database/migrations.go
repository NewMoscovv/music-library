package database

import (
	"Music-library/internal/models"
	myLogger "Music-library/pkg/logger"
)

func Migrate() {
	err := DB.AutoMigrate(&models.Song{})
	if err != nil {
		myLogger.Fatal("Ошибка миграции базы данных", map[string]interface{}{"error": err.Error()})
	}
	myLogger.Info("Миграция БД успешно завершена", nil)
}
