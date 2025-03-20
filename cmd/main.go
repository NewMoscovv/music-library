package main

import (
	"Music-library/config"
	"Music-library/internal/database"
	"Music-library/internal/gateway/postgres"
	"Music-library/internal/routes"
	myLogger "Music-library/pkg/logger"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {

	// инициируем конфиг
	cfg := config.Init()

	// инициируем логирование
	myLogger.Init(cfg.LogLevel)

	// подключаемся к базе данных
	database.Init(cfg)

	// запускаем миграции
	database.Migrate()

	// инициируем gateway
	songGateway := postgres.NewPgSongGateway(database.DB)

	// создаем роутер
	router := gin.Default()

	// передаем Gateway в обработчик
	routes.SetupRoutes(router, songGateway)

	// запускаем сервак
	port := fmt.Sprintf(":%s", cfg.AppPort)
	myLogger.Info("Сервер запущен на порту"+cfg.AppPort, nil)
	router.Run(port)
}
