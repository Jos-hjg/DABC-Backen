package controllers

import (
	"dabc/email"
	"dabc/models"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
)

func SendCode(ctx *gin.Context) {
	mail := models.Mail{}
	if err := ctx.ShouldBind(&mail); err != nil {
		if _, ok := err.(validator.ValidationErrors); ok {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"code": http.StatusBadRequest,
				"msg":  models.ValidateError(err.(validator.ValidationErrors), models.MailValidation),
			})
		}
		return
	}
	//TODO: send email
	code := email.GetRand()
	to := []string{mail.To}
	email.SendEmail(to, mail.Subject, "<p>验证吗：</p><h3>"+string(code[0])+"</h3>")
	//TODO: redis record?
}

func GetCode2Verify(ctx *gin.Context) {
	rd := email.GetRand()
	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  rd,
	})
}
