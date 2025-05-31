package router

import (
	"gin-api-template/handlers"
	"gin-api-template/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// 使用日志中间件
	r.Use(middlewares.LoggingMiddleware())

	// API 路由组
	api := r.Group("/api")
	{
		// 健康检查接口
		api.GET("/health", handlers.HealthCheckHandler)
	}

	return r
}
