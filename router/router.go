package router

import (
	ac "dabc/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"Authorization"}

	r.Use(cors.New(config))
	//TODO: any router?
	users := r.Group("/users")
	users.POST("/login", ac.UserLogin)
	users.POST("/auth", ac.CheckAuth)
	users.GET("/find", ac.FindUser)
	users.PUT("/update", ac.UpdateUser)
	//announce := r.Group("/announcement")
	//announce.GET("/", ac.Announce)
	r.GET("/announcement", ac.Announce)

	return r
}
