package v1

import (
	"github.com/gin-gonic/gin"
	"mountain/middleware"
	"mountain/model"
	"mountain/utils/errmessage"
	"net/http"
)

func Login(c *gin.Context) {
	var data model.User
	c.ShouldBindJSON(&data)
	var token string
	code := model.CheckLogin(data.Username, data.Password)
	if code == errmessage.SUCCESS {
		token, code = middleware.Settoken(data.Username)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmessage.Geterrmessage(code),
		"token":   token,
	})
}
