package crud

import (
	"gin-api-template/infra"
	"gin-api-template/models"
	"gin-api-template/utils"

	"gorm.io/gorm"
)

// GetUserByID 根据ID获取用户 - 返回 ORM 模型
func GetUserByID(id uint) (*models.User, error) {
	utils.LogInfo("CRUD: Getting user by ID")

	var user models.User
	err := infra.GetDB().Where("id = ? AND yn = ?", id, true).First(&user).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.LogInfo("CRUD: User not found")
			return nil, nil // 用户不存在，返回 nil 而不是错误
		}
		utils.LogError("CRUD: Failed to get user: " + err.Error())
		return nil, err
	}

	return &user, nil
}

// GetUserByPhone 根据手机号获取用户
func GetUserByPhone(phone string) (*models.User, error) {
	utils.LogInfo("CRUD: Getting user by phone")

	var user models.User
	err := infra.GetDB().Where("phone = ? AND yn = ?", phone, true).First(&user).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		utils.LogError("CRUD: Failed to get user by phone: " + err.Error())
		return nil, err
	}

	return &user, nil
}

// CreateUser 创建用户
func CreateUser(user *models.User) error {
	utils.LogInfo("CRUD: Creating user")

	err := infra.GetDB().Create(user).Error
	if err != nil {
		utils.LogError("CRUD: Failed to create user: " + err.Error())
		return err
	}

	return nil
}

// UpdateUser 更新用户
func UpdateUser(user *models.User) error {
	utils.LogInfo("CRUD: Updating user")

	err := infra.GetDB().Save(user).Error
	if err != nil {
		utils.LogError("CRUD: Failed to update user: " + err.Error())
		return err
	}

	return nil
}

// DeleteUser 软删除用户 (设置 yn = false)
func DeleteUser(id uint) error {
	utils.LogInfo("CRUD: Soft deleting user")

	err := infra.GetDB().Model(&models.User{}).Where("id = ?", id).Update("yn", false).Error
	if err != nil {
		utils.LogError("CRUD: Failed to delete user: " + err.Error())
		return err
	}

	return nil
}
