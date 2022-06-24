package main

import (
	"dabc/config"
	"dabc/contract"
	"dabc/database"
	"dabc/router"
	"log"
)

func main() {
	//读取配置文件
	config.InitCfg("./config/config.yaml")
	contract.Init()
	//初始化 DB
	if db, err := database.InitDB(); err != nil {
		log.Fatal(err)
	} else {
		defer db.Close()
	}

	//email.InitMail()
	//to := []string{"1287935492@qq.com"}
	//err := email.SendEmail(to, "test", "<h1>DABC Email Test</h1>\n<button>button</button>")
	//if err != nil {
	//	log.Println(err)
	//}
	//if rdb, err := redis.InitRedis(); err != nil {
	//	log.Fatal(err)
	//} else {
	//	defer rdb.Close()
	//}

	r := router.InitRouter()
	r.Run(":" + config.C.Router.Port)
}
