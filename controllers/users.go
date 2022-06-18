package controllers

import (
	"dabc/config"
	"dabc/database"
	"dabc/models"
	"dabc/signature"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
	"log"
	"net/http"
	"strings"
	"time"
)

/*用户登录*/
func UserLogin(ctx *gin.Context) {
	isuser := models.Users{}
	data := &models.UserLogin{}
	if err := ctx.ShouldBind(&data); err != nil {
		if _, ok := err.(validator.ValidationErrors); ok {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"code": http.StatusBadRequest,
				"msg":  models.ValidateError(err.(validator.ValidationErrors), models.LoginValidation),
			})
		}
		return
	}
	log.Println(data)
	if !signature.VerifySig(data.Address, data.Signature, []byte(data.Msg)) {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"code": http.StatusUnauthorized,
			"msg":  "签名信息不匹配",
		})
		return
	}

	database.Mysql.Where("address = ?", data.Address).First(&isuser)

	if (isuser.Address == ""){
		if err := database.Mysql.Create(&models.Users{
			Address: data.Address,
			Nickname: data.Address,
			Email: "",
		}).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"code": http.StatusInternalServerError,
				"msg": err,
			})
			return
		}
	}


	claims := &jwt.StandardClaims{
		Audience:  data.Address,
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
	ctx.JSON(http.StatusOK, gin.H{
		"code":  http.StatusOK,
		"token": sign,
		"msg":   "",
	})

}

/*
验证签名登录
Header附带Authorization签名参数
签名前缀必须以"dabc-t "开头
*/
func CheckAuth(ctx *gin.Context) {
	headers := ctx.GetHeader("Authorization")
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

	token, _ := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.C.Auth.SignKey), nil
	})
	if token == nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"code": http.StatusUnauthorized,
			"msg":  "token invalid",
		})
		return
	}
	if !token.Valid {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"code": http.StatusUnauthorized,
			"msg":  "token invalid",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "",
	})
}

func FindUser(ctx *gin.Context) {
	isuser := models.Users{}
	data := models.User{}
	if err := ctx.ShouldBindQuery(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg": err,
		})
		return
	}

	database.Mysql.Where("address = ?", data.Address).First(&isuser)

	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg": isuser,
	})

}

/*用户注册*/
//func CreateUser(ctx *gin.Context) {
//	NewUser := models.User{}
//	if err := ctx.ShouldBind(&NewUser); err != nil {
//		if _, ok := err.(validator.ValidationErrors); ok {
//			ctx.JSON(http.StatusBadRequest, gin.H{
//				"code": http.StatusBadRequest,
//				"msg":  models.ValidateError(err.(validator.ValidationErrors), models.CreateValidation),
//			})
//		}
//		return
//	}
//	//TODO: check signature to confirm address
//
//	NewUser.Password = fmt.Sprintf("%x", sha512.Sum512([]byte(NewUser.Password+NewUser.UserName+config.C.Auth.Salt)))
//	user := models.Users{
//		Username: NewUser.UserName,
//		Nickname: NewUser.NickName,
//		Password: NewUser.Password,
//		Phone:    NewUser.Phone,
//		Address:  NewUser.Address,
//	}
//	if err := database.Mysql.Create(&user).Error; err != nil {
//		ctx.JSON(http.StatusOK, gin.H{
//			"code": http.StatusOK,
//			"msg":  err,
//		})
//		return
//	}
//	ctx.JSON(http.StatusOK, gin.H{
//		"code": http.StatusOK,
//		"msg":  "",
//	})
//}

/*用户信息更改*/
func UpdateUser(ctx *gin.Context) {
	isuser := models.Users{}
	data := &models.UserLogin{}
	if err := ctx.ShouldBindQuery(data); err != nil {
		if _, ok := err.(validator.ValidationErrors); ok {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"code": http.StatusBadRequest,
				"msg":  models.ValidateError(err.(validator.ValidationErrors), models.LoginValidation),
			})
		}
		return
	}
	//TODO: signature is needed!
	if !signature.VerifySig(data.Address, data.Signature, []byte(data.Msg)) {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"code": http.StatusUnauthorized,
			"msg":  "签名信息不匹配",
		})
		return
	}

	database.Mysql.Where("address = ?", data.Address).First(&isuser)

	NewUser := models.User{}
	if err := ctx.ShouldBind(&NewUser); err != nil {
		if _, ok := err.(validator.ValidationErrors); ok {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"code": http.StatusBadRequest,
				"msg":  models.ValidateError(err.(validator.ValidationErrors), models.Updatevalidation),
			})
		}
		return
	}

	user := models.Users{
		Nickname: NewUser.NickName,
		Address:  NewUser.Address,
		Email:    NewUser.Email,
	}
	if err := database.Mysql.Model(&isuser).Update(user).Error; err != nil {
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
