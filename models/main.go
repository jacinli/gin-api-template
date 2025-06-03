package models

// GetAllModels 返回所有需要迁移的模型
// 添加新模型时，只需要在这里加一行即可
func GetAllModels() []interface{} {
	return []interface{}{
		&User{},
		// 新模型只需要在这里添加一行
		// &Product{},
		// &Order{},
	}
}
