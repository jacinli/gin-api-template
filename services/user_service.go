package services

import (
	"errors"

	"gin-api-template/crud"
	"gin-api-template/models"
	"gin-api-template/utils"
)

// UserResponse 用户响应结构 - 序列化后的数据
type UserResponse struct {
	ID         uint   `json:"id"`
	Phone      string `json:"phone"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
}

// GetUserByID 根据ID获取用户 - 返回序列化数据
func GetUserByID(id uint) (*UserResponse, error) {
	utils.LogInfo("Service: Getting user by ID")

	// 调用 CRUD 层
	user, err := crud.GetUserByID(id)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, nil // 用户不存在
	}

	// 转换为响应格式
	response := &UserResponse{
		ID:         user.Id,
		Phone:      user.Phone,
		CreateTime: user.CreateTime.Format("2006-01-02 15:04:05"),
		UpdateTime: user.UpdateTime.Format("2006-01-02 15:04:05"),
	}

	return response, nil
}

// GetUserByPhone 根据手机号获取用户
func GetUserByPhone(phone string) (*UserResponse, error) {
	utils.LogInfo("Service: Getting user by phone")

	user, err := crud.GetUserByPhone(phone)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, nil
	}

	response := &UserResponse{
		ID:         user.Id,
		Phone:      user.Phone,
		CreateTime: user.CreateTime.Format("2006-01-02 15:04:05"),
		UpdateTime: user.UpdateTime.Format("2006-01-02 15:04:05"),
	}

	return response, nil
}

// CreateUserRequest 创建用户请求结构
type CreateUserRequest struct {
	Phone    string `json:"phone" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
}

// CreateUser 创建用户
func CreateUser(req *CreateUserRequest) (*UserResponse, error) {
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

	// 返回创建的用户信息
	return GetUserByID(user.Id)
}

// GetUserList 获取用户列表
func GetUserList(page, pageSize int) ([]UserResponse, error) {
	utils.LogInfo("Service: Getting user list")

	offset := (page - 1) * pageSize
	users, err := crud.GetAllUsers(pageSize, offset)
	if err != nil {
		return nil, err
	}

	// 转换为响应格式
	var responses []UserResponse
	for _, user := range users {
		response := UserResponse{
			ID:         user.Id,
			Phone:      user.Phone,
			CreateTime: user.CreateTime.Format("2006-01-02 15:04:05"),
			UpdateTime: user.UpdateTime.Format("2006-01-02 15:04:05"),
		}
		responses = append(responses, response)
	}

	return responses, nil
}
