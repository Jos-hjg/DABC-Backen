package controllers

import (
	"dabc/config"
	"dabc/models"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
	"log"
	"net/http"
	"strings"
	"time"
)

func UserLogin(ctx *gin.Context) {
	data := &models.UserLogin{}
	if err := ctx.ShouldBind(data); err != nil {
		if _, ok := err.(validator.ValidationErrors); ok {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": models.LoginError(err.(validator.ValidationErrors)),
			})
		}

		return
	}

	claims := &jwt.StandardClaims{
		Audience:  data.Username,
		ExpiresAt: time.Now().Add(10 * 60 * time.Second).Unix(),
		Issuer:    "DABC-Backen",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	sign, err := token.SignedString([]byte(config.C.Auth.SignKey))
	if err != nil {
		ctx.JSON(http.StatusNetworkAuthenticationRequired, gin.H{
			"code": http.StatusNetworkAuthenticationRequired,
			"msg":  "token 签名失败",
		})
	}
	//log.Println(sign)
	ctx.JSON(http.StatusOK, gin.H{
		"token": sign,
		"msg":   "",
	})

}

func CheckAuth(ctx *gin.Context) {
	headers := ctx.GetHeader("Authorization")
	log.Println(headers)
	if haspre := strings.HasPrefix(headers, "Bearer "); !haspre {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"code": http.StatusUnauthorized,
			"msg":  "error",
		})
		return
	}
	if len(headers) <= 7 {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"code": http.StatusUnauthorized,
			"msg":  "have no authorization",
		})
		return
	}
	tokenStr := headers[7:]

	//token := sha256.Sum256([]byte("I am a token"))
	token, _ := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.C.Auth.SignKey), nil
	})
	if !token.Valid {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"code": http.StatusUnauthorized,
			"msg":  "token is not available",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"msg":    "",
	})
}
