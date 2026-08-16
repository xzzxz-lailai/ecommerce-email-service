package router

import (
	"email-service/handler/email"

	"github.com/gin-gonic/gin"
)

// EmailRoutes 邮箱路由
func EmailRoutes(r *gin.RouterGroup) {
	emailGroup := r.Group("/email")
	{
		emailGroup.POST("/code/send", email.SendEmailCode)     // 发送邮箱验证码
		emailGroup.POST("/code/verify", email.VerifyEmailCode) // 验证邮箱验证码
	}
}
