package controllers

import (
	"dabc/config"
	"dabc/contract"
	"dabc/database"
	"dabc/email"
	"dabc/models"
	"dabc/signature"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
	"log"
	"net/http"
	"strconv"
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
			"msg":  "signature error",
		})
		return
	}

	database.Mysql.Where("address = ?", data.Address).First(&isuser)

	if isuser.Address == "" {
		if err := database.Mysql.Create(&models.Users{
			Address:  data.Address,
			Nickname: data.Address,
			Email:    "",
			Pledged:  false,
		}).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"code": http.StatusInternalServerError,
				"msg":  err,
			})
			return
		}
	} else {
		//TODO: judge pledged
		//using contract.Client
		minter, err := contract.Client.Minters(nil, common.HexToAddress(data.Address))
		if err != nil {
			to := []string{"1287935492@qq.com"}
			email.SendEmail(to, "server error", "<h3>Error message:</h3><body>"+err.Error()+"</body>")
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"code": http.StatusInternalServerError,
				"msg":  err,
			})
			return
		}
		tbLength, _ := strconv.Atoi(minter.Tblength.String())
		if tbLength != 0 && !isuser.Pledged {
			//TODO: database's pledged
			database.Mysql.Model(&isuser).Update(models.Users{Pledged: true})
		}
	}

	claims := &jwt.StandardClaims{
		Audience:  data.Address,
		ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
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
		ctx.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "type error",
		})
		return
	}
	if len(headers) <= 7 {
		ctx.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "have no authorization",
		})
		return
	}
	tokenStr := headers[7:]

	token, _ := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.C.Auth.SignKey), nil
	})
	if token == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "token invalid",
		})
		return
	}
	if !token.Valid {
		ctx.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
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
		if _, ok := err.(validator.ValidationErrors); ok {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"code": http.StatusBadRequest,
				"msg":  models.ValidateError(err.(validator.ValidationErrors), models.LoginValidation),
			})
		}
		return
	}

	database.Mysql.Where("address = ?", data.Address).First(&isuser)

	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg": struct {
			Address  string
			NickName string
		}{
			Address:  isuser.Address,
			NickName: isuser.Nickname,
		},
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
	//TODO: before this require, need to send an eamil to verify
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
			"msg":  "signature error",
		})
		return
	}

	database.Mysql.Where("address = ?", data.Address).First(&isuser)

	if isuser.Address == "" {
		ctx.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "user is not exist",
		})
	}

	NewUser := models.Update{}

	if err := ctx.ShouldBind(&NewUser); err != nil {
		if _, ok := err.(validator.ValidationErrors); ok {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"code": http.StatusBadRequest,
				"msg":  models.ValidateError(err.(validator.ValidationErrors), models.UpdateValidation),
			})
		}
		return
	}
	if NewUser.Email != isuser.Email {
		//TODO: verify new email
	}

	user := models.Users{
		Nickname: NewUser.NickName,
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
