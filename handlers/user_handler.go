package handlers

import (
	"fmt"
	"strconv"

	"gin-api-template/constants"
	"gin-api-template/services"
	"gin-api-template/utils"

	"github.com/gin-gonic/gin"
)

// GetUserHandler 根据ID获取用户
func GetUserHandler(c *gin.Context) {
	utils.LogInfo("Handler: Get user called")

	idStr := c.Param("id")
	if idStr == "" {
		constants.BadRequest(c, "用户ID不能为空")
		return
	}

	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		constants.BadRequest(c, "用户ID格式错误")
		return
	}

	user, err := services.GetUserByID(uint(id))
	if err != nil {
		constants.InternalServerError(c, "获取用户失败")
		return
	}

	if user == nil {
		constants.NotFound(c, "用户不存在")
		return
	}

	constants.Success(c, user)
}

type GetUserRequest struct {
	ID uint `json:"id" binding:"required"`
}

func GetUserHandlerByPost(c *gin.Context) {
	utils.LogInfo("Handler: Get user called")

	var req GetUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		constants.BadRequest(c, "参数错误")
		return
	}
	userIDRaw, exists := c.Get("user_id")
	if !exists {
		constants.BadRequest(c, "用户ID不存在")
		return
	}
	fmt.Println("userIDRaw", userIDRaw)
	userID, ok := userIDRaw.(uint) // 或者根据你 claims.UserID 的类型来断言，比如 uint、int64
	if !ok {
		constants.BadRequest(c, "用户ID类型错误")
		return
	}

	utils.LogInfo(fmt.Sprintf("当前用户ID: %d", userID))
	user, err := services.GetUserByID(req.ID)
	if err != nil {
		constants.InternalServerError(c, "获取用户失败")
		return
	}

	constants.Success(c, user)
}
