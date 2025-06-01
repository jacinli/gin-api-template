package middlewares

import (
	"gin-api-template/utils"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	allowedOrigins := []string{}

	if utils.AppConfig != nil && len(utils.AppConfig.CORSAllowedOrigins) > 0 {
		allowedOrigins = utils.AppConfig.CORSAllowedOrigins
	} else if utils.IsDevelopment() {
		allowedOrigins = []string{
			"http://localhost:3000",
			"http://localhost:8000",
			"http://127.0.0.1:3000",
		}
	}

	config := cors.Config{
		AllowOrigins:     allowedOrigins,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "X-Request-ID"},
		ExposeHeaders:    []string{"X-Request-ID"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}

	return cors.New(config)
}
