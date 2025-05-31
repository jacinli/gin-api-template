package constants

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// APIResponse 统一响应结构
type APIResponse struct {
	Code      int         `json:"code"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data,omitempty"`
	RequestID string      `json:"request_id"`
	Timestamp int64       `json:"timestamp"`
}

const RequestIDKey = "X-Request-ID"

// getRequestID 从 Context 获取 Request ID
func getRequestID(c *gin.Context) string {
	if requestID, exists := c.Get(RequestIDKey); exists {
		return requestID.(string)
	}
	return "unknown"
}

// Success 成功响应
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, APIResponse{
		Code:      200,
		Message:   "success",
		Data:      data,
		RequestID: getRequestID(c),
		Timestamp: time.Now().Unix(),
	})
}

// SuccessWithMessage 带自定义消息的成功响应
func SuccessWithMessage(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, APIResponse{
		Code:      200,
		Message:   message,
		Data:      data,
		RequestID: getRequestID(c),
		Timestamp: time.Now().Unix(),
	})
}

// Error 错误响应
func Error(c *gin.Context, code int, message string) {
	c.JSON(code, APIResponse{
		Code:      code,
		Message:   message,
		RequestID: getRequestID(c),
		Timestamp: time.Now().Unix(),
	})
}

// BadRequest 400 错误
func BadRequest(c *gin.Context, message string) {
	Error(c, http.StatusBadRequest, message)
}

// NotFound 404 错误
func NotFound(c *gin.Context, message string) {
	Error(c, http.StatusNotFound, message)
}

// InternalServerError 500 错误
func InternalServerError(c *gin.Context, message string) {
	Error(c, http.StatusInternalServerError, message)
}
