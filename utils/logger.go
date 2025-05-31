package utils

import (
	"context"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/gin-gonic/gin"
)

var (
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
)

const RequestIDKey = "X-Request-ID"

func init() {
	InfoLogger = log.New(os.Stdout, "[INFO] ", log.Ldate|log.Ltime)
	ErrorLogger = log.New(os.Stderr, "[ERROR] ", log.Ldate|log.Ltime)
}

// getRequestIDFromContext 从 context 获取 Request ID
func getRequestIDFromContext(ctx interface{}) string {
	if ctx == nil {
		return ""
	}

	// 如果是 *gin.Context
	if ginCtx, ok := ctx.(*gin.Context); ok {
		if requestID, exists := ginCtx.Get(RequestIDKey); exists {
			return requestID.(string)
		}
	}

	// 如果是 context.Context
	if stdCtx, ok := ctx.(context.Context); ok {
		if requestID := stdCtx.Value(RequestIDKey); requestID != nil {
			return requestID.(string)
		}
	}

	return ""
}

// LogInfo 记录信息日志
func LogInfo(msg string) {
	_, file, line, ok := runtime.Caller(1)

	if ok {
		filename := filepath.Base(file)
		InfoLogger.Printf("%s:%d: %s", filename, line, msg)
	} else {
		InfoLogger.Println(msg)
	}
}

// LogInfoWithContext 带 context 的信息日志 (自动提取 Request ID)
func LogInfoWithContext(ctx interface{}, msg string) {
	_, file, line, ok := runtime.Caller(1)
	requestID := getRequestIDFromContext(ctx)

	if ok {
		filename := filepath.Base(file)
		if requestID != "" {
			InfoLogger.Printf("%s:%d [%s]: %s", filename, line, requestID, msg)
		} else {
			InfoLogger.Printf("%s:%d: %s", filename, line, msg)
		}
	} else {
		if requestID != "" {
			InfoLogger.Printf("[%s]: %s", requestID, msg)
		} else {
			InfoLogger.Println(msg)
		}
	}
}

// LogError 记录错误日志
func LogError(msg string) {
	_, file, line, ok := runtime.Caller(1)

	if ok {
		filename := filepath.Base(file)
		ErrorLogger.Printf("%s:%d: %s", filename, line, msg)
	} else {
		ErrorLogger.Println(msg)
	}
}

// LogErrorWithContext 带 context 的错误日志 (自动提取 Request ID)
func LogErrorWithContext(ctx interface{}, msg string) {
	_, file, line, ok := runtime.Caller(1)
	requestID := getRequestIDFromContext(ctx)

	if ok {
		filename := filepath.Base(file)
		if requestID != "" {
			ErrorLogger.Printf("%s:%d [%s]: %s", filename, line, requestID, msg)
		} else {
			ErrorLogger.Printf("%s:%d: %s", filename, line, msg)
		}
	} else {
		if requestID != "" {
			ErrorLogger.Printf("[%s]: %s", requestID, msg)
		} else {
			ErrorLogger.Println(msg)
		}
	}
}

// 删除 logWithRequestID 函数，直接在上面的函数中处理
