package utils

import (
	"log"
	"os"
	"path/filepath"
	"runtime"
)

var (
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
)

func init() {
	InfoLogger = log.New(os.Stdout, "[INFO] ", log.Ldate|log.Ltime)
	ErrorLogger = log.New(os.Stderr, "[ERROR] ", log.Ldate|log.Ltime)
}

// LogInfo 记录信息日志 - 自动包含当前 goroutine 的 Request ID
func LogInfo(msg string) {
	_, file, line, ok := runtime.Caller(1)
	requestID := GetRequestID() // 自动获取当前 goroutine 的 Request ID

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

// LogError 记录错误日志 - 自动包含当前 goroutine 的 Request ID
func LogError(msg string) {
	_, file, line, ok := runtime.Caller(1)
	requestID := GetRequestID() // 自动获取当前 goroutine 的 Request ID

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
