package middlewares

import (
	"bytes"
	"fmt"
	"io"
	"time"

	"gin-api-template/utils"

	"github.com/gin-gonic/gin"
)

// LoggingMiddleware API请求响应日志中间件
func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		// 读取请求体
		var requestBody []byte
		if c.Request.Body != nil {
			requestBody, _ = io.ReadAll(c.Request.Body)
			// 重新设置请求体，因为读取后会被消耗
			c.Request.Body = io.NopCloser(bytes.NewBuffer(requestBody))
		}

		// 创建自定义的响应写入器来捕获响应
		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw

		// 记录请求信息 - 自动包含 Request ID
		utils.LogInfoWithContext(c, fmt.Sprintf("Request: %s %s | IP: %s | User-Agent: %s | Body: %s",
			c.Request.Method,
			c.Request.URL.Path,
			c.ClientIP(),
			c.Request.UserAgent(),
			string(requestBody)))

		// 处理请求
		c.Next()

		// 计算处理时间
		duration := time.Since(startTime)

		// 记录响应信息 - 自动包含 Request ID
		utils.LogInfoWithContext(c, fmt.Sprintf("Response: %s %s | Status: %d | Duration: %v | Response: %s",
			c.Request.Method,
			c.Request.URL.Path,
			c.Writer.Status(),
			duration,
			blw.body.String()))
	}
}

// bodyLogWriter 自定义响应写入器
type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}
