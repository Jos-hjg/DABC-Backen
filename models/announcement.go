package models

import "github.com/jinzhu/gorm"

type Announces struct {
	gorm.Model
	SubCN string `json:"subcn"`
	SubEn string `json:"suben"`
	CN string `json:"cn""`
	EN string `json:"en"`
}
