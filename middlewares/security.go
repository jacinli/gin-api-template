package middlewares

import "github.com/gin-gonic/gin"

func SecurityMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 设置安全头 (在处理请求之前)
		c.Header("X-Frame-Options", "DENY")
		c.Header("Content-Security-Policy", "frame-ancestors 'none'")
		c.Header("Referrer-Policy", "no-referrer")
		c.Header("Cross-Origin-Opener-Policy", "same-origin")
		c.Header("X-Content-Type-Options", "nosniff")
		c.Header("X-XSS-Protection", "1; mode=block")

		// HTTPS 才设置 HSTS
		if c.Request.TLS != nil || c.GetHeader("X-Forwarded-Proto") == "https" {
			c.Header("Strict-Transport-Security", "max-age=63072000; includeSubDomains; preload")
		}

		// 处理请求
		c.Next()
	}
}
