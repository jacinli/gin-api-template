package infra

import (
	"fmt"
	"time"

	"gin-api-template/utils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// InitPG 初始化 PostgreSQL 连接
func InitPG() {
	config := utils.AppConfig
	if config == nil {
		utils.LogError("Config not loaded")
		return
	}

	// 构建数据库连接字符串
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		config.DBHost,
		config.DBUser,
		config.DBPassword,
		config.DBName,
		config.DBPort,
	)

	// 配置 GORM
	gormConfig := &gorm.Config{}

	// 根据环境设置日志级别
	if utils.IsDevelopment() {
		gormConfig.Logger = logger.Default.LogMode(logger.Info)
		utils.LogInfo("PostgreSQL debug mode enabled")
	} else {
		gormConfig.Logger = logger.Default.LogMode(logger.Silent)
	}

	// 连接数据库
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), gormConfig)
	if err != nil {
		utils.LogError("Failed to connect to PostgreSQL: " + err.Error())
		panic(err)
	}

	// 配置连接池
	sqlDB, err := DB.DB()
	if err != nil {
		utils.LogError("Failed to get PostgreSQL instance: " + err.Error())
		panic(err)
	}

	// 设置连接池参数
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	utils.LogInfo("PostgreSQL connected successfully")
}

// GetDB 获取数据库实例
func GetDB() *gorm.DB {
	return DB
}

// ClosePG 关闭数据库连接
func ClosePG() {
	if DB != nil {
		sqlDB, err := DB.DB()
		if err != nil {
			utils.LogError("Failed to get PostgreSQL instance for closing: " + err.Error())
			return
		}

		if err := sqlDB.Close(); err != nil {
			utils.LogError("Failed to close PostgreSQL: " + err.Error())
		} else {
			utils.LogInfo("PostgreSQL connection closed")
		}
	}
}
