package router

import (
	"gin-api-template/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	// 使用 gin.New() 而不是 gin.Default()，避免默认的日志中间件
	r := gin.New()

	// 添加 Recovery 中间件（防止 panic 导致服务崩溃）
	r.Use(gin.Recovery())
	// 使用安全头中间件 (应该在最前面)
	r.Use(middlewares.SecurityMiddleware())
	// 使用 CORS 中间件
	r.Use(middlewares.CORSMiddleware())

	// 使用 Request ID 中间件
	r.Use(middlewares.RequestIDMiddleware())

	// 使用我们自定义的日志中间件
	r.Use(middlewares.LoggingMiddleware())

	// 使用 JWT 认证中间件
	r.Use(middlewares.AuthMiddleware())

	// 注册各个模块的路由
	setupHealthRoutes(r)
	setupUserRoutes(r)
	setupWebSocketRoutes(r)
	// setupUserRoutes(r)
	// setupAuthRoutes(r)
	// setupProductRoutes(r)
	// 在这里添加更多路由模块...

	return r
}
