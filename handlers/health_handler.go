package handlers

import (
	"gin-api-template/constants"
	"gin-api-template/services"
	"gin-api-template/utils"

	"github.com/gin-gonic/gin"
)

// HealthCheckHandler 健康检查处理器
func HealthCheckHandler(c *gin.Context) {
	// 使用带 context 的日志，自动包含 Request ID
	utils.LogInfoWithContext(c, "Health check handler called")

	// 调用业务逻辑层
	result := services.GetHealthStatus()

	// 使用统一响应格式
	constants.Success(c, result)
}
