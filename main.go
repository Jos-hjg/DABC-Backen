package main

import (
	"dabc/config"
	"dabc/database"
	"dabc/router"
	"log"
)

func main() {
	//读取配置文件
	config.InitCfg("./config/config.yaml")
	//初始化 DB
	if db, err := database.InitDB(); err != nil {
		log.Fatal(err)
	} else {
		defer db.Close()
	}

	//if rdb, err := redis.InitRedis(); err != nil {
	//	log.Fatal(err)
	//} else {
	//	defer rdb.Close()
	//}

	r := router.InitRouter()
	r.Run(":" + config.C.Router.Port)
}
