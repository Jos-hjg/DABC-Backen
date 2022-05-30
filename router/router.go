package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.Use(cors.Default())
	//TODO: any router?
	//r.GET("/router", controller.methods)

	return r
}
