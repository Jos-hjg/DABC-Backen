package router

import (
	ac "dabc/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.Use(cors.Default())
	//TODO: any router?
	users := r.Group("/users")
	users.POST("/login", ac.UserLogin)
	users.POST("/auth", ac.CheckAuth)
	//users.POST("/create", ac.CreateUser)
	users.PUT("/update", ac.UpdateUser)
	return r
}
