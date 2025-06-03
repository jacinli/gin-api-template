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

	// 动态获取所有模型
	allModels := models.GetAllModels()
	err := DB.AutoMigrate(allModels...)

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

	// 动态获取所有模型并删除表
	allModels := models.GetAllModels()
	err := DB.Migrator().DropTable(allModels...)

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

	// 获取所有模型
	allModels := models.GetAllModels()

	for _, model := range allModels {
		if DB.Migrator().HasTable(model) {
			// 获取表名 (通过 TableName 方法或者类型名)
			if tableNamer, ok := model.(interface{ TableName() string }); ok {
				fmt.Printf("✅ Table exists: %s\n", tableNamer.TableName())
			} else {
				fmt.Printf("✅ Table exists: %v\n", model)
			}
		} else {
			if tableNamer, ok := model.(interface{ TableName() string }); ok {
				fmt.Printf("❌ Table missing: %s\n", tableNamer.TableName())
			} else {
				fmt.Printf("❌ Table missing: %v\n", model)
			}
		}
	}
}
