package controllers

import (
	"dabc/config"
	"dabc/contract"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetDABCContract(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "",
		"abi":  contract.DABCABI,
		"abr":  config.C.Chain.Contract.Address,
	})
}
