package tests

import (
	"testing"

	"gin-api-template/services"

	"github.com/stretchr/testify/assert"
)

// 测试纯业务逻辑 - 不依赖任何框架
func TestGetHealthStatus(t *testing.T) {
	result := services.GetHealthStatus()

	assert.Equal(t, "ok", result.Status)
	assert.Equal(t, "ping", result.Message)
}

// 测试业务逻辑的其他场景
func TestHealthResponseStructure(t *testing.T) {
	result := services.GetHealthStatus()

	// 验证返回结构
	assert.NotEmpty(t, result.Status)
	assert.NotEmpty(t, result.Message)
}
