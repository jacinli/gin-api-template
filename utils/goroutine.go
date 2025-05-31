package utils

import (
	"runtime"
	"strconv"
	"strings"
)

// getGoroutineID 获取当前 goroutine ID
// 注意：这是一个 hack 方法，在生产环境中谨慎使用
func getGoroutineID() int64 {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.ParseInt(idField, 10, 64)
	if err != nil {
		return 0
	}
	return id
}
