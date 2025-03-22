package main

import (
	"Music-library/config"
	_ "Music-library/docs"
	"Music-library/internal/gateway/postgres"
	"Music-library/internal/routes"
	database2 "Music-library/pkg/database"
	myLogger "Music-library/pkg/logger"
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Music Library API
// @version 1.0
// @description Приложение для онлайн-библиотеки музыки на RestAPI
// @host localhost:8080
// @BasePath /
func main() {

	// 1. Загружаем конфиг
	cfg := config.Init()

	// 2. Инициализируем логирование
	myLogger.Init(cfg.LogLevel)

	// 3. Подключаем к базе данных
	database2.Init(cfg)

	// 4. Запускаем миграции
	database2.Migrate()

	// 5. Инициализируем gateway
	songGateway := postgres.NewPgSongGateway(database2.DB)

	// 6. Создаем роутер
	router := gin.Default()

	// 7. Регистрируем маршруты
	routes.SetupRoutes(router, songGateway)

	// 8. Swagger UI
	router.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 9. Запускаем сервак
	port := fmt.Sprintf(":%s", cfg.AppPort)
	myLogger.Info("Сервер запущен на порту"+cfg.AppPort, nil)
	router.Run(port)
}
