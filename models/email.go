package models

type Mail struct {
	To string `json:"to" form:"to" binding:"required"`
	Subject string `json:"subject" form:"subject" binding:"required"`
}

type Mails struct {
	To []string `json:"to" form:"to" binding:"required"`
	Subject string `json:"subject" form:"subject"`
}

var MailValidation = ErrorType{
	"To": {
		"required": "缺少邮箱地址",
	},
	"Subject": {
		"required": "缺少标题",
	},
}