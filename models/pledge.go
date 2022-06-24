package models

import (
	"github.com/jinzhu/gorm"
)

type Pledge struct {
	gorm.Model
	Address string `json:"address" form:"address" binding:"required"`
	Amount string `json:"amount" form:"amount" binding:"required"`

}