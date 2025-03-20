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
	database.Init(cfg, logger)

	// запускаем миграции
	database.Migrate(logger)

	// вывод сообщения об успешном запуске
	fmt.Println("приложение запущено")

}
