package services

import (
	"errors"

	"gin-api-template/crud"
	"gin-api-template/models"
	"gin-api-template/utils"
)

// CreateUserRequest 创建用户请求结构
type CreateUserRequest struct {
	Phone    string `json:"phone" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
}

// GetUserByID 根据ID获取用户 - 直接返回模型 (已有 json 标签)
func GetUserByID(id uint) (*models.User, error) {
	utils.LogInfo("Service: Getting user by ID")

	// 调用 CRUD 层
	user, err := crud.GetUserByID(id)
	if err != nil {
		return nil, err
	}

	// 直接返回模型，models.User 已经有 json 标签了
	return user, nil
}

// GetUserByPhone 根据手机号获取用户
func GetUserByPhone(phone string) (*models.User, error) {
	utils.LogInfo("Service: Getting user by phone")

	user, err := crud.GetUserByPhone(phone)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// CreateUser 创建用户
func CreateUser(req *CreateUserRequest) (*models.User, error) {
	utils.LogInfo("Service: Creating user")

	// 检查用户是否已存在
	existingUser, err := crud.GetUserByPhone(req.Phone)
	if err != nil {
		return nil, err
	}

	if existingUser != nil {
		return nil, errors.New("用户已存在")
	}

	// 这里应该加密密码
	// hashedPassword := hashPassword(req.Password)

	// 创建用户模型
	user := &models.User{
		Phone:    req.Phone,
		Password: req.Password, // 实际应该是加密后的密码
	}

	err = crud.CreateUser(user)
	if err != nil {
		return nil, err
	}

	// 返回创建的用户信息 (Password 字段有 json:"-"，不会序列化)
	return user, nil
}
