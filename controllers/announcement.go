package controllers

import (
	"dabc/database"
	"dabc/models"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
)


func Announce(ctx *gin.Context)  {
	var total int
	announce := &[]models.Announces{}
	page := &models.Pages{}
	if err := ctx.ShouldBindQuery(&page); err != nil {
		if _, ok := err.(validator.ValidationErrors); ok {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"code": http.StatusBadRequest,
				"msg":  models.ValidateError(err.(validator.ValidationErrors), models.PagesValidation),
			})
		}
		return
	}
	db := database.Mysql.Model(models.Announces{})
	db.Count(&total)
	if err := db.Order("id DESC").Offset((page.Page - 1) * page.PageSize).Limit(page.PageSize).Find(&announce).Error; err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg": "数据查询异常",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg": "success",
		"data": map[string]interface{}{
			"data": announce,
			"total": total,
			"page": page.Page,
			"pageSize": page.PageSize,
		},
	})
}
