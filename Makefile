# Go 项目管理命令

# 运行应用
run:
	go run main.go

# 构建应用
build:
	go build -o main .

# 数据库迁移
migrate-up:
	go run cmd/migrate.go -action=up

# 回滚迁移
migrate-down:
	go run cmd/migrate.go -action=down

# 查看迁移状态
migrate-status:
	go run cmd/migrate.go -action=status

# 安装依赖
deps:
	go mod tidy
	go mod download

# 运行测试
test:
	go test ./tests/

# 运行测试并显示覆盖率
test-cover:
	go test -cover ./tests/

# 运行特定测试
test-user:
	go test ./tests/ -run TestGetUserByID

# 运行测试 - 详细输出
test-verbose:
	go test -v ./tests/

# 格式化代码
fmt:
	go fmt ./...

# 代码检查
lint:
	golangci-lint run

# Docker 构建
docker-build:
	docker build -t gin-api-template .

# Docker 运行
docker-run:
	docker-compose up -d

# Docker 停止
docker-down:
	docker-compose down

# 清理构建文件
clean:
	rm -f main
	go clean

.PHONY: run build migrate-up migrate-down migrate-status deps test fmt lint docker-build docker-run docker-down clean