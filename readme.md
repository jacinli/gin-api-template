
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


docker-compose build

docker-compose up -d
docker build -t gin-api-template .


安装 swagger
go install github.com/swaggo/swag/cmd/swag@latest

cursor 调试：
go install github.com/go-delve/delve/cmd/dlv@latest

安装 pg 本地测试
brew install postgresql
brew services start postgresql

psql postgres

-- 创建数据库
CREATE DATABASE ginapitest;

-- 创建用户
CREATE USER myuser WITH PASSWORD 'ginapitest';

-- 给用户赋权
GRANT ALL PRIVILEGES ON DATABASE ginapitest TO myuser;


开发流程
# 1. 安装依赖
make deps

# 2. 运行数据库迁移
make migrate-up

# 3. 查看迁移状态
make migrate-status

# 4. 启动应用
make run

# 5. 如果需要回滚
make migrate-down