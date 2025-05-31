package utils

import (
	"context"
	"sync"
)

const RequestIDKey = "X-Request-ID"

// goroutine-local storage 模拟
var (
	contextMap = make(map[int64]context.Context)
	contextMu  sync.RWMutex
)

// SetRequestContext 设置当前 goroutine 的 context
func SetRequestContext(ctx context.Context) {
	contextMu.Lock()
	defer contextMu.Unlock()

	goroutineID := getGoroutineID()
	contextMap[goroutineID] = ctx
}

// GetRequestContext 获取当前 goroutine 的 context
func GetRequestContext() context.Context {
	contextMu.RLock()
	defer contextMu.RUnlock()

	goroutineID := getGoroutineID()
	if ctx, exists := contextMap[goroutineID]; exists {
		return ctx
	}
	return context.Background()
}

// ClearRequestContext 清理当前 goroutine 的 context
func ClearRequestContext() {
	contextMu.Lock()
	defer contextMu.Unlock()

	goroutineID := getGoroutineID()
	delete(contextMap, goroutineID)
}

// GetRequestID 从当前 goroutine 的 context 获取 Request ID
func GetRequestID() string {
	ctx := GetRequestContext()
	if requestID := ctx.Value(RequestIDKey); requestID != nil {
		return requestID.(string)
	}
	return ""
}
