package model

const (
	CodeTypeRegister       = "register"
	CodeTypeForgetPassword = "forget_password"
)

// SendEmailCodeRequest 发送邮箱验证码请求
type SendEmailCodeRequest struct {
	Email    string `json:"email" binding:"required,email"`
	CodeType string `json:"code_type" binding:"required"`
}

// VerifyEmailCodeRequest 验证邮箱验证码请求
type VerifyEmailCodeRequest struct {
	Email    string `json:"email" binding:"required,email"`
	CodeType string `json:"code_type" binding:"required"`
	Code     string `json:"code" binding:"required,len=6"`
}
