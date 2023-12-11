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

// 删除传感器
func DeleteSensor(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code = dao.DeleteSensor(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmessage.Geterrmessage(code),
	})
}

// 编辑传感器
func EditSensor(c *gin.Context) {
	var data model.Sensor
	var OrgSensor model.Sensor
	id, _ := strconv.Atoi(c.Param("id"))
	OrgSensor, code = dao.GetSensorByID(id)
	if code != errmessage.SUCCESS {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  code,
			"message": errmessage.Geterrmessage(code),
		})
		return
	}
	_ = c.ShouldBindJSON(&data)
	//如果没有传入某个值，就用原来的值
	data.SensorSN = OrgSensor.SensorSN
	if data.Latitude == 0 {
		data.Latitude = OrgSensor.Latitude
	}
	if data.Longitude == 0 {
		data.Longitude = OrgSensor.Longitude
	}
	if data.Altitude == 0 {
		data.Altitude = OrgSensor.Altitude
	}
	if data.Address == "" {
		data.Address = OrgSensor.Address
	}
	if data.AcquisitionInterval == 0 {
		data.AcquisitionInterval = OrgSensor.AcquisitionInterval
	}
	if data.ArchiveStatus == 0 {
		data.ArchiveStatus = OrgSensor.ArchiveStatus
	}
	if data.IsWarning == 0 {
		data.IsWarning = OrgSensor.IsWarning
	}
	if data.WarningInterval == 0 {
		data.WarningInterval = OrgSensor.WarningInterval
	}
	if data.Commit == "" {
		data.Commit = OrgSensor.Commit
	}
	if data.ProjectID == 0 {
		data.ProjectID = OrgSensor.ProjectID
	}
	msg, IsLegal := myValidator.MyValidate(data)
	if IsLegal != errmessage.SUCCESS {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  IsLegal,
			"message": msg,
		})
		return
	}
	code = dao.EditSensor(id, &data)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmessage.Geterrmessage(code),
	})
}
