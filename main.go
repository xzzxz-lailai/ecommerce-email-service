package main

import (
	"email-service/config"
	"email-service/etcd"
	"email-service/router"
)

func main() {
	// 初始化配置（包含 Redis 初始化）
	cfg := config.InitConfig()

	// 初始化etcd
	if err := etcd.InitEtcd(); err != nil {
		panic(err)
	}
	defer etcd.CloseEtcd()          // 关闭etcd
	if err := etcd.RegisterService( // email-service服务注册
		config.Cfg.Etcd.ServerName, config.Cfg.Etcd.ServeAddress); err != nil {
		panic(err)
	}

	// 初始化路由
	r := router.Router()

	// 启动服务
	r.Run(cfg.Server.Port)
}
