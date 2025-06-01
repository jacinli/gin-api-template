package router

import (
	"gin-api-template/handlers"
	"gin-api-template/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	// 使用 gin.New() 而不是 gin.Default()，避免默认的日志中间件
	r := gin.New()

	// 添加 Recovery 中间件（防止 panic 导致服务崩溃）
	r.Use(gin.Recovery())
	// 使用 CORS 中间件 (必须在其他中间件之前)
	r.Use(middlewares.CORSMiddleware())
	// 使用 Request ID 中间件 (必须在日志中间件之前)
	r.Use(middlewares.RequestIDMiddleware())

	// 使用我们自定义的日志中间件
	r.Use(middlewares.LoggingMiddleware())

	// API 路由组
	api := r.Group("/api")
	{
		// 健康检查接口
		api.GET("/health", handlers.HealthCheckHandler)
	}

	return r
}
