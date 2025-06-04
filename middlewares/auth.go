package middlewares

import (
	"gin-api-template/constants"
	"gin-api-template/crud"
	"gin-api-template/utils"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware JWT 认证中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		utils.LogInfo("认证中间件开始验证")

		// 从请求头获取授权信息
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			utils.LogError("未提供授权头")
			constants.Error(c, 401, "未提供授权令牌")
			c.Abort()
			return
		}

		// 提取令牌
		tokenString, err := utils.ExtractTokenFromHeader(authHeader)
		if err != nil {
			utils.LogError("提取令牌失败: " + err.Error())
			constants.Error(c, 401, "无效的授权头格式")
			c.Abort()
			return
		}

		// 验证令牌
		claims, err := utils.ValidateToken(tokenString)
		if err != nil {
			utils.LogError("令牌验证失败: " + err.Error())
			constants.Error(c, 401, "无效或过期的令牌")
			c.Abort()
			return
		}

		// 验证用户是否存在
		user, err := crud.GetUserByID(claims.UserID)
		if err != nil {
			utils.LogError("获取用户信息失败: " + err.Error())
			constants.Error(c, 500, "内部服务器错误")
			c.Abort()
			return
		}

		if user == nil {
			utils.LogError("用户不存在")
			constants.Error(c, 401, "用户不存在")
			c.Abort()
			return
		}

		// 将用户信息存储到上下文中
		c.Set("user", user)
		c.Set("user_id", claims.UserID)
		c.Set("user_phone", claims.Phone)

		utils.LogInfo("认证成功，用户ID: " + string(rune(claims.UserID)))

		// 继续处理请求
		c.Next()
	}
}

// OptionalAuthMiddleware 可选认证中间件 - 如果有令牌则验证，没有则继续
func OptionalAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			// 没有提供令牌，继续执行
			c.Next()
			return
		}

		// 提取并验证令牌
		tokenString, err := utils.ExtractTokenFromHeader(authHeader)
		if err != nil {
			// 令牌格式错误，继续执行但不设置用户信息
			c.Next()
			return
		}

		claims, err := utils.ValidateToken(tokenString)
		if err != nil {
			// 令牌无效，继续执行但不设置用户信息
			c.Next()
			return
		}

		// 验证用户是否存在
		user, err := crud.GetUserByID(claims.UserID)
		if err != nil || user == nil {
			// 用户不存在，继续执行但不设置用户信息
			c.Next()
			return
		}

		// 设置用户信息到上下文
		c.Set("user", user)
		c.Set("user_id", claims.UserID)
		c.Set("user_phone", claims.Phone)

		c.Next()
	}
}
