package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	AppPort string
	DBHost  string
	DBPort  string
	DBUser  string
	DBPass  string
	DBName  string
	APIUrl  string
}

func Init() (*Config, error) {

	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}

	return &Config{
		AppPort: getEnv("APP_PORT", "8080"),
		DBHost:  getEnv("DB_HOST", "localhost"),
		DBPort:  getEnv("DB_PORT", "5432"),
		DBUser:  getEnv("DB_USER", "postgres"),
		DBPass:  getEnv("DB_PASS", "postgres"),
		DBName:  getEnv("DB_NAME", "postgres"),
		APIUrl:  getEnv("API_URL", "http://localhost:8080"),
	}, nil

}

func getEnv(key, defaultValut string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValut
}
