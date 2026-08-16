package template

import (
	"email-service/model"
	"errors"
	"fmt"
)

// BuildEmailContent 根据验证码类型构建邮件主题和正文内容。
func BuildEmailContent(codeType, code string) (string, string, error) {
	switch codeType {
	case model.CodeTypeRegister:
		return "注册验证码", fmt.Sprintf("您的注册验证码是：%s", code), nil
	case model.CodeTypeForgetPassword:
		return "找回密码验证码", fmt.Sprintf("您的找回密码验证码是：%s", code), nil
	default:
		return "", "", errors.New("验证码类型错误")
	}
}
