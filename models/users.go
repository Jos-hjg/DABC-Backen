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
	Nickname string `json:"nickname" form:"nickname" gorm:"unique"`
	Email    string `json:"email" form:"email"`
	Pledged  bool   `json:"pledged" form:"pledged"`
}

type User struct {
	Address   string `json:"address" form:"address" binding:"required"`
	NickName  string `json:"nickname" form:"nickname"`
	Signature string `json:"signature" form:"signature"`
	Email     string `json:"email" form:"email"`
}

type Find struct {
	Address string `json:"address" form:"address" binding:"required"`
}

type Update struct {
	NickName    string `json:"nickname" form:"nickname" binding:"required"`
	Email       string `json:"email" form:"email"`
	EamilVerify string `json:"emailVerify" form:"emailVerify"`
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

var UpdateValidation = ErrorType{
	"Signature": {
		"required": "签名必须",
	},
}
