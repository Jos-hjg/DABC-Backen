package models

import (
	"github.com/jinzhu/gorm"
)

type UserLogin struct {
	Address   string `json:"address" form:"address" binding:"required"`
	Msg       string `json:"msg" form:"msg" binding:"required"`
	Signature string `json:"signature" form:"signature" binding:"required"`
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
	Email     string `json:"email" form:"email" binding:"required"`
}

var LoginValidation = ErrorType{
	"Signature": {
		"required": "签名必须",
	},
	"Address": {
		"required": "钱包地址必须",
	},
	"Msg": {
		"required": "签名信息必须",
	},
}

var Updatevalidation = ErrorType{
	"Signature": {
		"required": "签名必须",
	},
	"Address": {
		"required": "钱包地址必须",
	},
	"Email": {
		"required": "邮箱必须",
	},
}
