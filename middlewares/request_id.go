package middlewares

import (
	"context"

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

		// 设置到标准 context.Context 中，供全局使用
		ctx := context.WithValue(c.Request.Context(), RequestIDKey, requestID)
		c.Request = c.Request.WithContext(ctx)

		// 设置响应头
		c.Header(RequestIDKey, requestID)

		// 继续处理请求
		c.Next()
	}
}
