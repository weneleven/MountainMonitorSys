package v1

import (
	"mountain/internal/dao"
	errmessage "mountain/pkg/errcode"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 获取warn
func GetDisplace(c *gin.Context) {
	datas, total := dao.GetDisplace()
	code = errmessage.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmessage.Geterrmessage(code),
		"data":    datas,
		"total":   total,
	})
}

// 通过设备号获取采集数据
func GetDisplaceBySN(c *gin.Context) {
	sn := c.Query("sn")
	data, total, code := dao.GetDisplaceBySensorSN(sn)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmessage.Geterrmessage(code),
		"data":    data,
		"total":   total,
	})
}
