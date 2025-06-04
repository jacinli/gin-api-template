package main

import (
	"gin-api-template/infra"
	"gin-api-template/router"
	"gin-api-template/utils"
)

// @title           Gin API Template
// @version         1.0
// @description     这是一个使用 Gin 框架的 API 模板

// @host      localhost:8080
// @BasePath  /api

func main() {
	// 1. 加载配置
	config := utils.LoadConfig()

	// 2. 初始化 PostgreSQL
	infra.InitPG()
	defer infra.ClosePG()

	// 3. 运行数据库迁移
	// infra.RunPGMigrations()

	// 4. 初始化路由
	r := router.SetupRouter()

	// 5. 启动服务器
	utils.LogInfo("Server starting on port " + config.ServerPort)
	if err := r.Run(":" + config.ServerPort); err != nil {
		utils.LogError("Failed to start server: " + err.Error())
	}
}
