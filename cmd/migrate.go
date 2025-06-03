package main

import (
	"flag"
	"fmt"
	"os"

	"gin-api-template/infra"
	"gin-api-template/utils"
)

func main() {
	// 命令行参数
	var action string
	flag.StringVar(&action, "action", "", "Migration action: up, down, status")
	flag.Parse()

	if action == "" {
		fmt.Println("Usage: go run cmd/migrate.go -action=<up|down|status>")
		fmt.Println("Examples:")
		fmt.Println("  go run cmd/migrate.go -action=up      # 运行迁移")
		fmt.Println("  go run cmd/migrate.go -action=down    # 回滚迁移")
		fmt.Println("  go run cmd/migrate.go -action=status  # 查看迁移状态")
		os.Exit(1)
	}

	// 加载配置
	utils.LoadConfig()

	// 初始化数据库连接
	infra.InitPG()
	defer infra.ClosePG()

	switch action {
	case "up":
		fmt.Println("Running migrations...")
		infra.RunPGMigrations()
		fmt.Println("✅ Migrations completed successfully")

	case "down":
		fmt.Println("Rolling back migrations...")
		infra.RollbackPGMigrations()
		fmt.Println("✅ Rollback completed successfully")

	case "status":
		fmt.Println("Checking migration status...")
		infra.CheckPGMigrationStatus()

	default:
		fmt.Printf("Unknown action: %s\n", action)
		os.Exit(1)
	}
}
