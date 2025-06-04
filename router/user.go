package router

import (
	"gin-api-template/handlers"

	"github.com/gin-gonic/gin"
)

func setupUserRoutes(r *gin.Engine) {
	users := r.Group("/api/users")
	{
		users.GET("/:id", handlers.GetUserHandler) // 只要这一个
		users.POST("/post_by_id", handlers.GetUserHandlerByPost)
	}
}
