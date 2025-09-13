#!/bin/bash

# 脚本用于更新Docker容器并执行数据库迁移

echo "开始更新Docker容器并执行数据库迁移..."

# 项目根目录
PROJECT_ROOT="$(dirname "$(dirname "$(dirname "$0"))")")
DOCKER_COMPOSE_FILE="$PROJECT_ROOT/deploy/docker-compose/docker-compose.yml"

# 检查docker-compose文件是否存在
if [ ! -f "$DOCKER_COMPOSE_FILE" ]; then
  echo "错误: docker-compose文件 $DOCKER_COMPOSE_FILE 不存在"
  exit 1
fi

# 停止并重新启动Docker容器
echo "停止并重新启动Docker容器..."
cd "$PROJECT_ROOT"
docker-compose -f "$DOCKER_COMPOSE_FILE" down
docker-compose -f "$DOCKER_COMPOSE_FILE" up -d

# 等待MySQL容器启动完成
echo "等待MySQL容器启动完成..."
sleep 20

# 执行迁移脚本
echo "执行数据库迁移..."
bash "$PROJECT_ROOT/server/scripts/run_migrations.sh"

echo "更新和迁移完成!"