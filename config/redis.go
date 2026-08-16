package config

import (
	"email-service/global"
	"fmt"
	"github.com/redis/go-redis/v9"
)

func InitRedis() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", Cfg.Redis.Host, Cfg.Redis.Port),
		Password: Cfg.Redis.Pass, // 空字符串 = 无密码
		DB:       Cfg.Redis.DB,
	})
	global.RDB = rdb
	fmt.Println("✅ Redis 连接成功！")
}
