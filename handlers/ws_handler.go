package handlers

import (
	"net/http"

	"gin-api-template/constants"
	"gin-api-template/services"
	"gin-api-template/utils"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// WebSocket 升级器
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// 允许所有来源（生产环境需要更严格的检查）
		return true
	},
}

// WebSocketHandler WebSocket 处理器
func WebSocketHandler(c *gin.Context) {
	utils.LogInfo("WebSocket 连接请求")

	// 升级 HTTP 连接为 WebSocket
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		utils.LogError("WebSocket 升级失败: " + err.Error())
		constants.Error(c, 500, "WebSocket 升级失败")
		return
	}

	defer func() {
		conn.Close()
		utils.LogInfo("WebSocket 连接关闭")
	}()

	// 发送欢迎消息
	welcomeMsg := services.GetWelcomeMessage()
	conn.WriteJSON(welcomeMsg)

	// 启动心跳服务（类似 Python 的迭代器）
	heartbeatChan := services.StartHeartbeat()

	// 启动心跳发送 goroutine
	go func() {
		for heartbeat := range heartbeatChan {
			if err := conn.WriteJSON(heartbeat); err != nil {
				utils.LogError("发送心跳消息失败: " + err.Error())
				return
			}
		}
	}()

	// 监听客户端消息（主循环，类似 Python 的 while True）
	for {
		var msg services.WebSocketMessage
		err := conn.ReadJSON(&msg)
		if err != nil {
			utils.LogError("读取消息失败: " + err.Error())
			break
		}

		// 调用 services 层处理业务逻辑
		response := services.ProcessMessage(msg)

		// 发送回复
		if err := conn.WriteJSON(response); err != nil {
			utils.LogError("发送回复失败: " + err.Error())
			break
		}
	}
}
