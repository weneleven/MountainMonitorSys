package v1

import (
	"github.com/gin-gonic/gin"
	"mountain/model"
	"mountain/utils/errmessage"
	"mountain/utils/myValidator"
	"net/http"
	"strconv"
)

var code int

// 增加用户
func AddUser(c *gin.Context) {
	var data model.User
	_ = c.ShouldBindJSON(&data) //ShouldBindWith使用指定的绑定引擎绑定传递的结构指针
	var msg string
	msg, code = myValidator.MyValidate(data)
	if code != errmessage.SUCCESS {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  code,
			"message": msg,
		})
		return
	}
	code = model.CheckUserName(data.Username)
	if code == errmessage.SUCCESS {
		model.CreatUser(&data)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,                           //状态码
		"data":    data,                           //数据
		"message": errmessage.Geterrmessage(code), //状态信息
	})
}

// 查询用户链表
func GetUsers(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	data, total := model.GetUsers(pageSize, pageNum)
	code = errmessage.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status":  code,                           //状态码
		"data":    data,                           //数据
		"total":   total,                          //总数
		"message": errmessage.Geterrmessage(code), //状态信息
	})
}

// 编辑用户
func EditUsers(c *gin.Context) {
	var data model.User
	id, _ := strconv.Atoi(c.Param("id"))
	code = model.CheckUserID(id)
	_ = c.ShouldBindJSON(&data) //ShouldBindWith使用指定的绑定引擎绑定传递的结构指针
	if code == errmessage.SUCCESS {
		code = model.CheckUserName(data.Username)
		if code == errmessage.SUCCESS {
			code = model.EditUser(id, &data)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,                           //状态码
		"message": errmessage.Geterrmessage(code), //状态信息
	})
}

// 删除用户
func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code = model.DeleteUser(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,                           //状态码
		"message": errmessage.Geterrmessage(code), //状态信息
	})
}
