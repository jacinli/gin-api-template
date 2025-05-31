air 配置

go mod init gin-api-template

项目结构：
├── handlers/          # HTTP 处理层 (依赖 Gin)
├── services/          # 业务逻辑层 (框架无关)
├── models/            # 数据模型
├── router/            # 路由配置
├── middlewares/       # 中间件
├── utils/             # 工具函数
└── tests/             # 测试


go mod tidy
go mod download