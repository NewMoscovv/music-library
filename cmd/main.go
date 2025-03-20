package main

import (
	"Music-library/config"
	"Music-library/internal/database"
	myLogger "Music-library/pkg/logger"
	"fmt"
)

func main() {

	// инициируем логирование
	logger := myLogger.Init()

	// инициируем конфиг
	cfg, err := config.Init()
	if err != nil {
		logger.Err.Fatalf("Error loading config: %v", err)
	}

	// подключаемся к базе данных
	db, err := database.Init(cfg)
	if err != nil {
		logger.Err.Fatalf("Error connecting to database: %v", err)
	}

	fmt.Println("Connected to database", db)
}
