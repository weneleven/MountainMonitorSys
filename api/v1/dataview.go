package v1

import (
	"github.com/gin-gonic/gin"
	"mountain/internal/dao"
	errmessage "mountain/pkg/errcode"
	"net/http"
)

// 获取采集数据
func GetSensorData(c *gin.Context) {
	datas , total := dao.GetSensorData()
	code = errmessage.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmessage.Geterrmessage(code),
		"data":    datas,
		"total":   total,
	})
}
