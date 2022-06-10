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
	Address  string `json:"address" form:"address" gorm:"unique"`
	Nickname string `json:"nickname" form:"nickname"`
	Email    string `json:"email" form:"email"`
}

type User struct {
	Address   string `json:"address" form:"address" binding:"required"`
	NickName  string `json:"nickname" form:"nickname"`
	Signature string `json:"signature" form:"signature" binding:"required"`
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
	"Signature": {
		"required": "签名必须",
	},
	"Address": {
		"required": "钱包地址必须",
	},
}
