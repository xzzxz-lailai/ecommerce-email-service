package email

import (
	"net/http"

	"email-service/model"
	"email-service/pkg"
	email "email-service/service/email"

	"github.com/gin-gonic/gin"
)

// SendEmailCode 发送邮箱验证码
func SendEmailCode(c *gin.Context) {
	var req model.SendEmailCodeRequest

	// 接收前端 JSON
	if err := c.ShouldBindJSON(&req); err != nil {
		pkg.Error(c, http.StatusBadRequest, "参数错误")
		return
	}

	// 调用 Service
	if err := email.SendEmailCode(c.Request.Context(), &req); err != nil {
		pkg.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	// 返回结果
	pkg.Success(c, "验证码发送成功", nil)
}

// VerifyEmailCode 验证邮箱验证码
func VerifyEmailCode(c *gin.Context) {
	var req model.VerifyEmailCodeRequest

	// 接收前端 JSON
	if err := c.ShouldBindJSON(&req); err != nil {
		pkg.Error(c, http.StatusBadRequest, "参数错误")
		return
	}

	// 调用 Service
	if err := email.VerifyEmailCode(c.Request.Context(), &req); err != nil {
		pkg.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	// 返回结果
	pkg.Success(c, "验证码验证成功", nil)
}
