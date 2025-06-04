package handlers

import (
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
