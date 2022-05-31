package database

import (
	"dabc/config"
	"dabc/controllers"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func InitDB() (*gorm.DB, error) {
	DSN := config.C.Mysql.User + ":" + config.C.Mysql.Passwd + "@tcp(" + config.C.Mysql.Host + ":" + config.C.Mysql.Port + ")/" + config.C.Mysql.Database + "?charset=utf8&parseTime=true"
	db, err := gorm.Open("mysql", DSN)
	if err != nil {
		return nil, err
	}
	controllers.Mysql = db
	return db, nil
}
