package router

import (
	"gin-api-template/handlers"

	"github.com/gin-gonic/gin"
)

func setupHealthRoutes(r *gin.Engine) {
	health := r.Group("/api")
	{
		health.GET("/health", handlers.HealthCheckHandler)
		// health.GET("/deep", handlers.DeepHealthCheckHandler) // 深度健康检查
	}
}
