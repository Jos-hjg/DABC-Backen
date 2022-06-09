package models

import (
	"github.com/jinzhu/gorm"
)

type UserLogin struct {
	Username string `json:"username" binding:"required" form:"username"`
	Password string `json:"password" binding:"required" form:"password"`
}

type Users struct {
	gorm.Model
	Username string `json:"username" binding:"required" form:"username" gorm:"unique"`
	Password string `json:"_" binding:"required,gte=8" form:"password"`
	Nickname string `json:"nickname" form:"nickname"`
	Address  string `json:"address" form:"address" gorm:"unique"`
	Phone    string `json:"phone" form:"phone"`
}

type User struct {
	UserName        string `json:"username" binding:"required" form:"username"`
	Password        string `json:"password" binding:"required,gte=8" form:"password"`
	PasswordConfirm string `json:"password_confirm" binding:"required,eqfield=Password" form:"password_confirm"`
	NickName        string `json:"nickname" form:"nickname"`
	Address         string `json:"address" form:"address" binding:"required"`
	Phone           string `json:"phone" form:"phone"`
	Signature       string `json:"signature" form:"signature" binding:"required"`
}

var LoginValidation = ErrorType{
	"Username": {
		"required": "用户名称必须",
	},
	"Password": {
		"required": "密码必须",
	},
}

var CreateValidation = ErrorType{
	"UserName": {
		"required": "用户名称必须",
	},
	"Password": {
		"required": "密码必须",
		"gte":      "密码长度至少8位",
	},
	"PasswordConfirm": {
		"required": "确认密码必须",
		"eqfield":  "密码不一致",
	},
	"Signature": {
		"required": "签名必须",
	},
	"Address": {
		"required": "钱包地址必须",
	},
}
