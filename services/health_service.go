package services

import "gin-api-template/utils"

// HealthResponse 健康检查响应结构
type HealthResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// GetHealthStatus 获取健康状态 - 纯业务逻辑
func GetHealthStatus() HealthResponse {
	utils.LogInfo("Health check service called")

	// 这里可以添加实际的健康检查逻辑
	// 比如检查数据库连接、缓存、外部服务等

	return HealthResponse{
		Status:  "ok",
		Message: "ping",
	}
}
