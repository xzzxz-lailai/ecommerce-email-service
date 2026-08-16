package email

import (
	"context"
	"email-service/model"
	"email-service/pkg"
	"email-service/repo"
	"email-service/smtp"
	emailtemplate "email-service/template"
)

// SendEmailCode 发送邮箱验证码
func SendEmailCode(ctx context.Context, req *model.SendEmailCodeRequest) error {
	// 1. 校验验证码类型
	if err := pkg.ValidateCodeType(req.CodeType); err != nil {
		return err
	}

	// 2. 生成 6 位验证码
	code := pkg.GenerateVerificationCode()

	// 3. 保存验证码到 Redis
	if err := repo.SaveCodeToRedis(ctx, req.Email, req.CodeType, code); err != nil {
		return err
	}

	// 4. 根据验证码类型选择邮件模板
	subject, body, err := emailtemplate.BuildEmailContent(req.CodeType, code)
	if err != nil {
		_ = repo.DeleteCodeFromRedis(ctx, req.Email, req.CodeType)
		return err
	}

	// 5. 发送邮件
	if err := sendEmail(ctx, req.Email, subject, body); err != nil {
		// 如果邮件发送失败,就删除 redis 里的验证码
		_ = repo.DeleteCodeFromRedis(ctx, req.Email, req.CodeType)
		return err
	}

	return nil
}

// VerifyEmailCode 验证邮箱验证码
func VerifyEmailCode(ctx context.Context, req *model.VerifyEmailCodeRequest) error {
	// 1. 校验验证码类型
	if err := pkg.ValidateCodeType(req.CodeType); err != nil {
		return err
	}
	// 2. 校验验证码是否正确
	if err := repo.VerifyCodeFromRedis(ctx, req.Email, req.CodeType, req.Code); err != nil {
		return err
	}
	// 3. 验证成功后删除验证码,避免重复使用
	return repo.DeleteCodeFromRedis(ctx, req.Email, req.CodeType)
}

// sendEmail 发送邮件
func sendEmail(ctx context.Context, to, subject, body string) error {
	// 调用smtp发送
	if err := smtp.Send(to, subject, body); err != nil {
		return err
	}

	return nil
}
