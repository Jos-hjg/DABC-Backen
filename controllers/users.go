package controllers

import (
	"crypto/sha512"
	"dabc/config"
	"dabc/database"
	"dabc/models"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
	"log"
	"net/http"
	"strings"
	"time"
)

func UserLogin(ctx *gin.Context) {
	salt := "the salt"
	user := models.Users{}
	data := &models.UserLogin{}
	if err := ctx.ShouldBind(data); err != nil {
		if _, ok := err.(validator.ValidationErrors); ok {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"code": http.StatusBadRequest,
				"msg":  models.ValidateError(err.(validator.ValidationErrors), models.LoginValidation),
			})
		}
		return
	}
	database.Mysql.Where("username = ?", data.Username).First(&user)
	fmt.Println(user)
	pwd := fmt.Sprintf("%x", sha512.Sum512([]byte(data.Password+salt)))
	fmt.Println(pwd)
	if pwd != user.Password {
		ctx.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "账户不存在或密码错误",
		})
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
		ctx.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "token 签名失败",
		})
	}
	//log.Println(sign)
	ctx.JSON(http.StatusOK, gin.H{
		"code":  http.StatusOK,
		"token": sign,
		"msg":   "",
	})

}

func CheckAuth(ctx *gin.Context) {
	headers := ctx.GetHeader("Authorization")
	log.Println(headers)
	if haspre := strings.HasPrefix(headers, "dabc-t "); !haspre {
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
		"code": http.StatusOK,
		"msg":  "",
	})
}
func CreateUser(ctx *gin.Context) {
	salt := "the salt"
	NewUser := models.User{}
	if err := ctx.ShouldBind(&NewUser); err != nil {
		if _, ok := err.(validator.ValidationErrors); ok {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"code": http.StatusBadRequest,
				"msg":  models.ValidateError(err.(validator.ValidationErrors), models.CreateValidation),
			})
		}
		return
	}
	//TODO: check signature to confirm address

	NewUser.Password = fmt.Sprintf("%x", sha512.Sum512([]byte(NewUser.Password+salt)))
	user := models.Users{
		Username: NewUser.UserName,
		Nickname: NewUser.NickName,
		Password: NewUser.Password,
		Phone:    NewUser.Phone,
		Address:  NewUser.Address,
	}
	if err := database.Mysql.Create(&user).Error; err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  err,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "",
	})
}

func ChangeUser(ctx *gin.Context) {
	headers := ctx.GetHeader("Authorization")
	log.Println(headers)
	if haspre := strings.HasPrefix(headers, "dabc-t "); !haspre {
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
	salt := "the salt"
	NewUser := models.User{}
	if err := ctx.ShouldBind(&NewUser); err != nil {
		if _, ok := err.(validator.ValidationErrors); ok {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"code": http.StatusBadRequest,
				"msg":  models.ValidateError(err.(validator.ValidationErrors), models.CreateValidation),
			})
		}
		return
	}
	//TODO: check signature to confirm address

	NewUser.Password = fmt.Sprintf("%x", sha512.Sum512([]byte(NewUser.Password+salt)))
	user := models.Users{
		Username: NewUser.UserName,
		Nickname: NewUser.NickName,
		Password: NewUser.Password,
		Phone:    NewUser.Phone,
		Address:  NewUser.Address,
	}

	if err := database.Mysql.Update(&user).Error; err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  err,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "",
	})
}
