package main

import (
	"dabc/config"
	"dabc/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

func main() {
	config.InitCfg("./config/config.yaml")
	//初始化数据库连接
	DSN := config.C.Mysql.User + ":" + config.C.Mysql.Passwd + "@tcp(" + config.C.Mysql.Host + ":" + config.C.Mysql.Port + ")/" + config.C.Mysql.Database + "?charset=utf8&parseTime=true"
	db, err := gorm.Open("mysql", DSN)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	//根据models模板实现自动生成数据表的sql语句进行数据表迁移
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.Users{}) //数据库存在则不会创建，但会被修改成对应的结构
	log.Println("Migration completed!")
}
