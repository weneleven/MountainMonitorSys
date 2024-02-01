package v1

import (
	"github.com/gin-gonic/gin"
	"mountain/internal/dao"
	errmessage "mountain/pkg/errcode"
	"net/http"
)

// 获取采集数据
func GetSensorData(c *gin.Context) {
	datas, total := dao.GetSensorData()
	code = errmessage.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmessage.Geterrmessage(code),
		"data":    datas,
		"total":   total,
	})
}

// 通过设备号获取采集数据
func GetSensorDataBySN(c *gin.Context) {
	sn := c.Query("sn")
	data, total, code := dao.GetSensorDataBySensorSN(sn)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmessage.Geterrmessage(code),
		"data":    data,
		"total":   total,
	})
}

//// 新增数据
//func AddSensorData(c *gin.Context) {
//	var data model.SensorData
//	_ = c.ShouldBindJSON(&data)
//	msg, code := myValidator.MyValidate(data)
//	if code != errmessage.SUCCESS {
//		c.JSON(http.StatusBadRequest, gin.H{
//			"status":  code,
//			"message": msg,
//		})
//		return
//	}
//	code = dao.AddSensorData(&data)
//	if code != errmessage.SUCCESS {
//		c.JSON(http.StatusOK, gin.H{
//			"status":  code,
//			"message": errmessage.Geterrmessage(code),
//		})
//		return
//	}
//}
