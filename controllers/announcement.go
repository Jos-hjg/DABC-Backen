package controllers

import (
	"dabc/database"
	"dabc/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Announce(ctx *gin.Context)  {
	announce := []models.Announces{}

	database.Mysql.Find(&announce)
	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg": announce,
	})
}
