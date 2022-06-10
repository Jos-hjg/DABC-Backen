package main

import (
	"dabc/config"
	"dabc/database"
	"dabc/router"
	"dabc/signature"
	"fmt"
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

	ok := signature.VerifySig("0x821b121D544cAb0a4F4d0ED2F1c2B14fAb4f969F",
		"0xd511d019fab160c0bd56a8f8019c526be57a81f080b261947bb151a5a3febd2862ad0e046a52ee75d7924385502b9d681ee2ec94b145e716566ab3447a8f11f31b",
		[]byte("asd"))
	fmt.Println("Verify: ", ok)

	//if rdb, err := redis.InitRedis(); err != nil {
	//	log.Fatal(err)
	//} else {
	//	defer rdb.Close()
	//}

	r := router.InitRouter()
	r.Run(":" + config.C.Router.Port)
}
