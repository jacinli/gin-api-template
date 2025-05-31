package handlers

import (
	"net/http"

	"gin-api-template/services"

	"github.com/gin-gonic/gin"
)

// HealthCheckHandler 健康检查处理器
func HealthCheckHandler(c *gin.Context) {
	// 调用业务逻辑层
	result := services.GetHealthStatus()

	// 返回 HTTP 响应
	c.JSON(http.StatusOK, result)
}
