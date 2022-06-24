package controllers

import (
	"dabc/email"
	"dabc/models"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
)

func SendCode(ctx *gin.Context)  {
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
	code := "123456"
	to := []string{mail.To}
	email.SendEmail(to, mail.Subject, "<h3>Code:</h3>" + code)
	//TODO: redis record?
}
