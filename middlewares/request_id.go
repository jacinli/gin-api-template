package middlewares

import (
	"context"

	"gin-api-template/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const RequestIDKey = "X-Request-ID"

// RequestIDMiddleware 请求ID中间件
func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 先检查请求头中是否已有 Request ID
		requestID := c.GetHeader(RequestIDKey)

		// 如果没有，则生成新的 UUID
		if requestID == "" {
			requestID = uuid.New().String()
		}

		// 将 Request ID 存储到 Context 中
		c.Set(RequestIDKey, requestID)

		// 创建带 Request ID 的 context
		ctx := context.WithValue(context.Background(), utils.RequestIDKey, requestID)

		// 设置到全局 context (当前 goroutine)
		utils.SetRequestContext(ctx)

		// 设置响应头
		c.Header(RequestIDKey, requestID)

		// 继续处理请求
		c.Next()

		// 请求结束后清理 context
		utils.ClearRequestContext()
	}
}
