package tests

import (
	"testing"

	"gin-api-template/infra"
	"gin-api-template/models"
	"gin-api-template/services"
	"gin-api-template/utils"
)

// setupTestDB 设置测试数据库
func setupTestDB() {
	// 加载配置
	utils.LoadConfig()

	// 初始化数据库连接
	infra.InitPG()

	// 运行迁移
	// infra.RunPGMigrations()
}

// cleanupTestDB 清理测试数据
func cleanupTestDB() {
	if infra.GetDB() != nil {
		// 清理测试数据
		infra.GetDB().Where("1 = 1").Delete(&models.User{})
	}
}

// TestGetUserByID 测试根据ID获取用户
func TestGetUserByID(t *testing.T) {
	setupTestDB()
	defer cleanupTestDB()

	// 1. 测试获取不存在的用户
	user, err := services.GetUserByID(999)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if user != nil {
		t.Errorf("Expected nil user, got %v", user)
	}

	// 2. 先创建一个测试用户
	createReq := &services.CreateUserRequest{
		Phone:    "13800138000",
		Password: "123456",
	}

	createdUser, err := services.CreateUser(createReq)
	if err != nil {
		t.Fatalf("Failed to create test user: %v", err)
	}

	// 3. 测试获取存在的用户
	user, err = services.GetUserByID(createdUser.Id)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if user == nil {
		t.Errorf("Expected user, got nil")
	}
	if user.Phone != "13800138000" {
		t.Errorf("Expected phone 13800138000, got %s", user.Phone)
	}
}

// TestCreateUser 测试创建用户
func TestCreateUser(t *testing.T) {
	setupTestDB()
	defer cleanupTestDB()

	// 1. 测试正常创建用户
	req := &services.CreateUserRequest{
		Phone:    "13800138001",
		Password: "123456",
	}

	user, err := services.CreateUser(req)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if user == nil {
		t.Errorf("Expected user, got nil")
	}
	if user.Phone != "13800138001" {
		t.Errorf("Expected phone 13800138001, got %s", user.Phone)
	}

	// 2. 测试创建重复用户（应该失败）
	duplicateUser, err := services.CreateUser(req)
	if err == nil {
		t.Errorf("Expected error for duplicate user, got nil")
	}
	if duplicateUser != nil {
		t.Errorf("Expected nil user for duplicate, got %v", duplicateUser)
	}
}

// TestGetUserByPhone 测试根据手机号获取用户
func TestGetUserByPhone(t *testing.T) {
	setupTestDB()
	defer cleanupTestDB()

	// 1. 测试获取不存在的用户
	user, err := services.GetUserByPhone("18888888888")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if user != nil {
		t.Errorf("Expected nil user, got %v", user)
	}

	// 2. 先创建一个测试用户
	createReq := &services.CreateUserRequest{
		Phone:    "13800138002",
		Password: "123456",
	}

	_, err = services.CreateUser(createReq)
	if err != nil {
		t.Fatalf("Failed to create test user: %v", err)
	}

	// 3. 测试获取存在的用户
	user, err = services.GetUserByPhone("13800138002")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if user == nil {
		t.Errorf("Expected user, got nil")
	}
	if user.Phone != "13800138002" {
		t.Errorf("Expected phone 13800138002, got %s", user.Phone)
	}
}

// TestCreateUserValidation 测试创建用户的参数验证
func TestCreateUserValidation(t *testing.T) {
	setupTestDB()
	defer cleanupTestDB()

	// 测试空手机号
	req := &services.CreateUserRequest{
		Phone:    "",
		Password: "123456",
	}

	user, err := services.CreateUser(req)
	if err == nil {
		t.Errorf("Expected error for empty phone, got nil")
	}
	if user != nil {
		t.Errorf("Expected nil user for invalid request, got %v", user)
	}
}
