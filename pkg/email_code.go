package pkg

import (
	"email-service/model"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// ValidateCodeType 校验验证码业务类型是否合法
func ValidateCodeType(codeType string) error {
	switch codeType {
	case model.CodeTypeRegister: // 注册
		return nil

	case model.CodeTypeForgetPassword: // 忘记密码
		return nil

	default:
		return errors.New("验证码类型错误")
	}
}

// GenerateVerificationCode 生成 6 位邮箱验证码
func GenerateVerificationCode() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return fmt.Sprintf("%06d", r.Intn(1000000))
}
