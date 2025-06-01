package main // 主函数需要这样写，不可以写 package gin-api-template

import (
	"gin-api-template/router"
	"gin-api-template/utils"
)

func main() {
	// 加载配置
	config := utils.LoadConfig()

	// 初始化路由
	r := router.SetupRouter()

	// 启动服务器
	utils.LogInfo("Starting " + config.AppName + " v" + config.AppVersion)
	utils.LogInfo("Server starting on " + utils.GetServerAddress())

	if err := r.Run(":" + config.ServerPort); err != nil {
		utils.LogError("Failed to start server: " + err.Error())
	}
}
