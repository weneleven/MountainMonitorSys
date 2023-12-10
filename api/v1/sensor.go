package v1

import (
	"github.com/gin-gonic/gin"
	"mountain/internal/dao"
	"mountain/internal/model"
	errmessage "mountain/pkg/errcode"
	"mountain/pkg/myValidator"
	"net/http"
	"strconv"
)

// 新增传感器
func AddSensor(c *gin.Context) {
	var data model.Sensor
	_ = c.ShouldBindJSON(&data)
	msg, code := myValidator.MyValidate(data)
	if code != errmessage.SUCCESS {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  code,
			"message": msg,
		})
		return
	}
	code = dao.CheckSensorSN(data.SensorSN)
	if code == errmessage.SUCCESS {
		code = dao.AddSensor(&data)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmessage.Geterrmessage(code),
	})
}

// 获取传感器列表
func GetSensors(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	sensors, total := dao.GetSensors(pageSize, pageNum)
	code = errmessage.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmessage.Geterrmessage(code),
		"data":    sensors,
		"total":   total,
	})
}
