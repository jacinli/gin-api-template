package utils

import (
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort string
	AppEnv     string

	// CORS 配置
	CORSAllowedOrigins []string

	// PostgreSQL 配置
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string
}

var AppConfig *Config

func LoadConfig() *Config {
	godotenv.Load()

	AppConfig = &Config{
		ServerPort:         getEnv("SERVER_PORT", "8080"),
		AppEnv:             getEnv("APP_ENV", "development"),
		CORSAllowedOrigins: getEnvAsSlice("CORS_ALLOWED_ORIGINS", []string{}),
		DBHost:             getEnv("DB_HOST", "localhost"),
		DBUser:             getEnv("DB_USER", "postgres"),
		DBPassword:         getEnv("DB_PASSWORD", "postgres"),
		DBName:             getEnv("DB_NAME", "postgres"),
		DBPort:             getEnv("DB_PORT", "5432"),
	}

	LogInfo("Config loaded, server port: " + AppConfig.ServerPort)
	return AppConfig
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvAsSlice(key string, defaultValue []string) []string {
	if value := os.Getenv(key); value != "" {
		parts := strings.Split(value, ",")
		result := make([]string, 0, len(parts))
		for _, part := range parts {
			if trimmed := strings.TrimSpace(part); trimmed != "" {
				result = append(result, trimmed)
			}
		}
		if len(result) > 0 {
			return result
		}
	}
	return defaultValue
}

func IsDevelopment() bool {
	return AppConfig != nil && AppConfig.AppEnv == "development"
}

func IsProduction() bool {
	return AppConfig != nil && AppConfig.AppEnv == "production"
}
