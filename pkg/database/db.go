package database

import (
	"Music-library/config"
	myLogger "Music-library/pkg/logger"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init(cfg *config.Config) {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.DBUser, cfg.DBPass, cfg.DBHost, cfg.DBPort, cfg.DBName,
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		myLogger.Fatal("Ошибка подключения к БД", map[string]interface{}{"error": err.Error()})
	}
	myLogger.Info("Подключение к БД успешно установлено", nil)
}
