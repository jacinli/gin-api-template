package utils

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config 应用配置结构
type Config struct {
	// 服务器配置
	ServerPort string
	ServerHost string

	// 应用配置
	AppName    string
	AppVersion string
	AppEnv     string

	// 日志配置
	LogLevel string

	// 数据库配置
	DBHost     string
	DBPort     string
	DBName     string
	DBUser     string
	DBPassword string

	// Redis 配置
	RedisHost     string
	RedisPort     string
	RedisPassword string

	// JWT 配置
	JWTSecret      string
	JWTExpireHours int
}

var AppConfig *Config

// LoadConfig 加载配置
func LoadConfig() *Config {
	// 加载 .env 文件
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using environment variables")
	}

	config := &Config{
		// 服务器配置
		ServerPort: getEnv("SERVER_PORT", "8080"),
		ServerHost: getEnv("SERVER_HOST", "0.0.0.0"),

		// 应用配置
		AppName:    getEnv("APP_NAME", "gin-api-template"),
		AppVersion: getEnv("APP_VERSION", "1.0.0"),
		AppEnv:     getEnv("APP_ENV", "development"),

		// 日志配置
		LogLevel: getEnv("LOG_LEVEL", "info"),

		// 数据库配置
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBName:     getEnv("DB_NAME", "gin_api"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "password"),

		// Redis 配置
		RedisHost:     getEnv("REDIS_HOST", "localhost"),
		RedisPort:     getEnv("REDIS_PORT", "6379"),
		RedisPassword: getEnv("REDIS_PASSWORD", ""),

		// JWT 配置
		JWTSecret:      getEnv("JWT_SECRET", "default-secret-change-in-production"),
		JWTExpireHours: getEnvAsInt("JWT_EXPIRE_HOURS", 24),
	}

	// 设置全局配置
	AppConfig = config

	LogInfo("Configuration loaded successfully")
	LogInfo("Server will start on " + config.ServerHost + ":" + config.ServerPort)
	LogInfo("Environment: " + config.AppEnv)

	return config
}

// getEnv 获取环境变量，如果不存在则返回默认值
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// getEnvAsInt 获取环境变量并转换为 int，如果不存在或转换失败则返回默认值
func getEnvAsInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}

// getEnvAsBool 获取环境变量并转换为 bool，如果不存在或转换失败则返回默认值
func getEnvAsBool(key string, defaultValue bool) bool {
	if value := os.Getenv(key); value != "" {
		if boolValue, err := strconv.ParseBool(value); err == nil {
			return boolValue
		}
	}
	return defaultValue
}

// IsProduction 判断是否为生产环境
func IsProduction() bool {
	return AppConfig != nil && AppConfig.AppEnv == "production"
}

// IsDevelopment 判断是否为开发环境
func IsDevelopment() bool {
	return AppConfig != nil && AppConfig.AppEnv == "development"
}

// GetServerAddress 获取服务器完整地址
func GetServerAddress() string {
	if AppConfig == nil {
		return ":8080"
	}
	return AppConfig.ServerHost + ":" + AppConfig.ServerPort
}
