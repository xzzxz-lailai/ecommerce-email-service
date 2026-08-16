package repo

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"email-service/global"

	"github.com/redis/go-redis/v9"
)

const emailCodeTTL = 5 * time.Minute

// SaveCodeToRedis 把验证码保存到 Redis,默认 5 分钟过期
func SaveCodeToRedis(ctx context.Context, email, codeType, code string) error {
	if global.RDB == nil {
		return errors.New("redis 未初始化")
	}

	return global.RDB.Set(ctx, buildEmailCodeKey(email, codeType), code, emailCodeTTL).Err()
}

// VerifyCodeFromRedis 从 Redis 取验证码，比对用户提交的 code
func VerifyCodeFromRedis(ctx context.Context, email, codeType, code string) error {
	if global.RDB == nil {
		return errors.New("redis 未初始化")
	}

	savedCode, err := global.RDB.Get(ctx, buildEmailCodeKey(email, codeType)).Result()
	if errors.Is(err, redis.Nil) {
		return errors.New("验证码不存在或已过期")
	}
	if err != nil {
		return err
	}
	if savedCode != code {
		return errors.New("验证码错误")
	}

	return nil
}

// DeleteCodeFromRedis 验证成功后删除验证码，避免重复使用
func DeleteCodeFromRedis(ctx context.Context, email, codeType string) error {
	if global.RDB == nil {
		return errors.New("redis 未初始化")
	}

	return global.RDB.Del(ctx, buildEmailCodeKey(email, codeType)).Err()
}

// buildEmailCodeKey 统一 Redis key
func buildEmailCodeKey(email, codeType string) string {
	return fmt.Sprintf("email_code:%s:%s", codeType, strings.ToLower(strings.TrimSpace(email)))
}
