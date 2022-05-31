package models

import (
	"github.com/jinzhu/gorm"
	"gopkg.in/go-playground/validator.v9"
)

type UserLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Users struct {
	gorm.Model
	UserName string `json:"username" binding:"required" form:"username"`
	Password string `json:"_" binding:"required" form:"password"`
	NickName string `json:"nickname" form:"nickname"`
	Address  string `json:"address" form:"address"`
	Phone    string `json:"phone" form:"phone"`
}

var LoginValidation = ErrorType{
	"Username": {
		"required": "用户名称必须",
	},
	"Password": {
		"required": "密码必须",
	},
}

func LoginError(errs validator.ValidationErrors) map[string]string {
	errMap := map[string]string{}
	for _, err := range errs {
		errMap[err.Field()] = LoginValidation[err.Field()][err.Tag()]
	}
	return errMap
}
