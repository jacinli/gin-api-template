package router

import (
	"gin-api-template/handlers"

	"github.com/gin-gonic/gin"
)

// setupWebSocketRoutes 设置 WebSocket 路由
func setupWebSocketRoutes(r *gin.Engine) {
	ws := r.Group("/api/ws")
	{
		ws.GET("/connect", handlers.WebSocketHandler)
	}
}
