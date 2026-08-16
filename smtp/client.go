package smtp

import (
	"fmt"

	"email-service/config"

	"github.com/go-gomail/gomail"
)

// Send 发送邮件
func Send(toEmail, subject, body string) error {
	// 创建邮件对象
	m := gomail.NewMessage()
	// 发件人
	m.SetHeader("From", m.FormatAddress(config.Cfg.EmailSmtp.Email, config.Cfg.EmailSmtp.SenderName))
	// 收件人
	m.SetHeader("To", toEmail)
	// 邮件主题
	m.SetHeader("Subject", subject)
	// 邮件正文
	m.SetBody("text/html", body)
	// 创建 SMTP 连接
	d := gomail.NewDialer(
		config.Cfg.EmailSmtp.Host,
		config.Cfg.EmailSmtp.Port,
		config.Cfg.EmailSmtp.Email,
		config.Cfg.EmailSmtp.AuthCode,
	)
	// 使用 SSL
	d.SSL = true
	// 发送邮件
	if err := d.DialAndSend(m); err != nil {
		return fmt.Errorf("邮件发送失败: %w", err)
	}

	return nil
}
