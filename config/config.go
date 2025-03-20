package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

type Config struct {
	AppPort  string
	DBHost   string
	DBPort   int
	DBUser   string
	DBPass   string
	DBName   string
	APIUrl   string
	LogLevel string
}

func Init() *Config {

	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Не удалось загрузить .env, используем переменные окружения")
	}

	dBPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatalf("Ошибка преобразования DB_PORT в int: %v", err)
	}

	return &Config{
		AppPort:  getEnv("APP_PORT", "8080"),
		DBHost:   getEnv("DB_HOST", "localhost"),
		DBUser:   getEnv("DB_USER", "postgres"),
		DBPort:   dBPort,
		DBPass:   getEnv("DB_PASSWORD", "postgres"),
		DBName:   getEnv("DB_NAME", "postgres"),
		APIUrl:   getEnv("API_URL", "http://localhost:8080"),
		LogLevel: getEnv("LOG_LEVEL", "debug"),
	}

}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
