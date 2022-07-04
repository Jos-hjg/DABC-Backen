package models

import "github.com/jinzhu/gorm"

type Announces struct {
	gorm.Model
	SubCN string `json:"subcn"`
	SubEn string `json:"suben"`
	CN string `json:"cn""`
	EN string `json:"en"`
}

type Pages struct {
	Page int `json:"page" form:"page" binding:"required"`
	PageSize int `json:"pagesize" form:"pagesize" binding:"required"`
}


var PagesValidation = ErrorType{
	"Page": {
		"required": "第几页必须",
	},
	"PageSize": {
		"required": "单页数量必须",
	},
}
