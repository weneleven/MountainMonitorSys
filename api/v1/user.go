package v1

import (
	"github.com/gin-gonic/gin"
	"mountain/internal/dao"
	"mountain/internal/model"
	"mountain/pkg/errcode"
	"mountain/pkg/myValidator"
	"net/http"
	"strconv"
)

var code int

// 增加用户
func AddUser(c *gin.Context) {
	var data model.User
	_ = c.ShouldBindJSON(&data) //ShouldBindWith使用指定的绑定引擎绑定传递的结构指针
	var msg string
	IsSuccess := true
	msg, code = myValidator.MyValidate(data)
	if code != errmessage.SUCCESS {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  code,
			"message": msg,
		})
		return
	}
	code = dao.CheckUserName(data.Username)
	if code != errmessage.SUCCESS {
		IsSuccess = false
	}
	code = dao.CheckUserPhone(data.Phone)
	if code != errmessage.SUCCESS {
		IsSuccess = false
	}
	if IsSuccess {
		code = dao.CreatUser(&data)
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
	data, total := dao.GetUsers(pageSize, pageNum)
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
	code = dao.CheckUserID(id)
	_ = c.ShouldBindJSON(&data) //ShouldBindWith使用指定的绑定引擎绑定传递的结构指针
	msg, Islegalcode := myValidator.MyValidate(data)
	if code != errmessage.SUCCESS {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  Islegalcode,
			"message": msg,
		})
		return
	}
	if code == errmessage.SUCCESS {
		code = dao.CheckUserName(data.Username)
		OrgUser, _ := dao.FindUser(id)
		if code == errmessage.SUCCESS || data.Username == OrgUser.Username {
			code = dao.EditUser(id, &data)
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
	code = dao.DeleteUser(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,                           //状态码
		"message": errmessage.Geterrmessage(code), //状态信息
	})
}
