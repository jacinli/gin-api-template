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

// LogInfo 记录信息日志
func LogInfo(msg string) {
	_, file, line, ok := runtime.Caller(1)
	if ok {
		// 使用 filepath.Base 获取文件名，简单多了
		filename := filepath.Base(file)
		InfoLogger.Printf("%s:%d: %s", filename, line, msg)
	} else {
		InfoLogger.Println(msg)
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
