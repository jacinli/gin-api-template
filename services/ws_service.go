package services

import (
	"time"

	"gin-api-template/utils"
)

// WebSocketMessage 简单的消息结构
type WebSocketMessage struct {
	Type    string `json:"type"`
	Content string `json:"content"`
	Time    string `json:"time"`
}

// MessageHandler 定义消息处理接口（类似 Python 的迭代器概念）
type MessageHandler interface {
	Send(msg WebSocketMessage) error
	Receive() (*WebSocketMessage, error)
	Close() error
}

// GetWelcomeMessage 获取欢迎消息
func GetWelcomeMessage() WebSocketMessage {
	return WebSocketMessage{
		Type:    "welcome",
		Content: "欢迎连接到 WebSocket 服务！",
		Time:    time.Now().Format("15:04:05"),
	}
}

// GetHeartbeatMessage 获取心跳消息
func GetHeartbeatMessage() WebSocketMessage {
	return WebSocketMessage{
		Type:    "heartbeat",
		Content: "服务器心跳消息",
		Time:    time.Now().Format("15:04:05"),
	}
}

// ProcessMessage 处理收到的消息（业务逻辑）
func ProcessMessage(msg WebSocketMessage) WebSocketMessage {
	utils.LogInfo("处理消息: " + msg.Content)

	// 模拟业务处理延迟
	time.Sleep(100 * time.Millisecond)

	// 返回处理后的消息
	return WebSocketMessage{
		Type:    "response",
		Content: "服务器收到: " + msg.Content,
		Time:    time.Now().Format("15:04:05"),
	}
}

// StartHeartbeat 启动心跳服务（返回一个通道，类似 Python 生成器）
func StartHeartbeat() <-chan WebSocketMessage {
	heartbeatChan := make(chan WebSocketMessage)

	go func() {
		defer close(heartbeatChan)
		ticker := time.NewTicker(3 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				heartbeatChan <- GetHeartbeatMessage()
			}
		}
	}()

	return heartbeatChan
}
