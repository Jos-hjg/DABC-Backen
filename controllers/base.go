package controllers

import (
	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
)

var Mysql *gorm.DB
var Redis *redis.Client
