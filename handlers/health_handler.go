package handlers

import (
	"gin-api-template/constants"
	"gin-api-template/services"
	"gin-api-template/utils"

	"github.com/gin-gonic/gin"
)

// HealthCheckHandler 健康检查处理器
func HealthCheckHandler(c *gin.Context) {
	// 直接使用 LogInfo，会自动包含 Request ID
	utils.LogInfo("Health check handler called")

	// 调用业务逻辑层
	result := services.GetHealthStatus()

	// 使用统一响应格式
	constants.Success(c, result)
}
