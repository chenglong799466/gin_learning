#!/bin/bash

# 创建标准Gin项目结构生成脚本
# 使用方法: ./create_project.sh your-project-name

set -e

PROJECT_NAME=$1

if [[ -z "$PROJECT_NAME" ]]; then
    echo "ERROR: 必须指定项目名称"
    echo "使用方法: ./create_project.sh your-project-name"
    exit 1
fi

# 创建基础目录结构
mkdir -p ${PROJECT_NAME}/{cmd,internal/api/{handler,middleware},internal/service,internal/repository,internal/model,internal/config,internal/pkg,pkg,test,scripts,configs,docs,web}

# 初始化Go模块
cd ${PROJECT_NAME}
go mod init ${PROJECT_NAME}

# 创建核心文件
cat > cmd/main.go <<EOF
package main

import (
	"${PROJECT_NAME}/internal/api"
)

func main() {
	router := api.NewRouter()
	router.Run(":8080")
}
EOF

# 创建Dockerfile
cat > Dockerfile <<EOF
FROM golang:1.20-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o server ./cmd/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/server .
COPY configs/ ./configs/
EXPOSE 8080
CMD ["./server"]
EOF

# 创建Makefile
cat > Makefile <<EOF
.PHONY: run build docker

run:
	go run cmd/main.go

build:
	go build -o bin/server cmd/main.go

docker:
	docker build -t ${PROJECT_NAME} .
EOF

# 创建示例配置文件
mkdir -p configs
cat > configs/config.yaml <<EOF
server:
  port: 8080
  env: development

database:
  dsn: "user:password@tcp(localhost:3306)/dbname?parseTime=true"
EOF

# 创建基础路由文件
cat > internal/api/router.go <<EOF
package api

import (
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	
	// 健康检查路由
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})
	
	return r
}
EOF

# 设置权限
chmod -R 755 scripts/

echo "项目 ${PROJECT_NAME} 创建完成！"
echo "目录结构："
tree -d -L 3 ${PROJECT_NAME}
