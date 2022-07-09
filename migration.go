package main

import (
	"dabc/config"
	"dabc/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"time"
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
	//db.DropTable(models.Users{})
	//db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.Announces{}) //数据库存在则不会创建，但会被修改成对应的结构
	for i := 0; i < 10; i++ {
		db.Create(&models.Announces{SubCN: time.Now().String() + " DABC 公告测试", SubEn: time.Now().String() + "DABC announcement Test", CN: "DABC 公告内容", EN: "DABC announcement content"})
		time.Sleep(3 * time.Second)
	}

	//db.Create(&models.Announces{SubCN: "关于 DABC ****的最新公告", SubEn: "About DABC's **** newly announcement", CN: "DABC 公告内容", EN: "DABC announcement content"})
	log.Println("Migration completed!")
}
