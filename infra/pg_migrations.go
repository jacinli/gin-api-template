package infra

import (
	"fmt"

	"gin-api-template/models"
	"gin-api-template/utils"
)

// RunPGMigrations 运行 PostgreSQL 数据库迁移
func RunPGMigrations() {
	if DB == nil {
		utils.LogError("PostgreSQL not initialized")
		return
	}

	utils.LogInfo("Starting PostgreSQL migration...")

	// 添加你的模型到这里
	err := DB.AutoMigrate(
		&models.User{},
		// 在这里添加更多模型...
	)

	if err != nil {
		utils.LogError("PostgreSQL migration failed: " + err.Error())
		panic(err)
	}

	utils.LogInfo("PostgreSQL migration completed successfully")
}

// RollbackPGMigrations 回滚数据库迁移
func RollbackPGMigrations() {
	if DB == nil {
		utils.LogError("PostgreSQL not initialized")
		return
	}

	utils.LogInfo("Starting PostgreSQL rollback...")

	// 删除表（谨慎操作！）
	err := DB.Migrator().DropTable(
		&models.User{},
		// 添加要删除的表...
	)

	if err != nil {
		utils.LogError("PostgreSQL rollback failed: " + err.Error())
		panic(err)
	}

	utils.LogInfo("PostgreSQL rollback completed successfully")
}

// CheckPGMigrationStatus 检查迁移状态
func CheckPGMigrationStatus() {
	if DB == nil {
		utils.LogError("PostgreSQL not initialized")
		return
	}

	fmt.Println("=== Migration Status ===")

	// 检查表是否存在
	tables := []interface{}{
		&models.User{},
		// 添加要检查的模型...
	}

	for _, table := range tables {
		if DB.Migrator().HasTable(table) {
			fmt.Printf("✅ Table exists: %s\n", DB.Statement.Schema.Table)
		} else {
			fmt.Printf("❌ Table missing: %s\n", DB.Statement.Schema.Table)
		}
	}
}
