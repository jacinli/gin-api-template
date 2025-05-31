package main // 主函数需要这样写，不可以写 package gin-api-template

import (
	"log"

	"gin-api-template/router"
)

func main() {
	// 初始化路由
	r := router.SetupRouter()

	// 启动服务器
	log.Println("Server starting on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
