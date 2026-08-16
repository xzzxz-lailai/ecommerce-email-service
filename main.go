package main

import (
	"email-service/config"
	"email-service/router"
)

func main() {
	// 初始化配置（包含 Redis 初始化）
	cfg := config.InitConfig()

	// 初始化路由
	r := router.Router()

	// 启动服务
	r.Run(cfg.Server.Port)
}
