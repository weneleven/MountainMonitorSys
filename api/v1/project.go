package v1

import (
	"fmt"
	"mountain/internal/dao"
	"mountain/internal/model"
	errmessage "mountain/pkg/errcode"
	"mountain/pkg/myValidator"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddProject(c *gin.Context) {
	var data model.Project
	_ = c.ShouldBindJSON(&data)
	msg, code := myValidator.MyValidate(data)
	if code != errmessage.SUCCESS {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  code,
			"message": msg,
		})
		return
	}
	dao.CreateProject(&data)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmessage.Geterrmessage(code),
	})
}

// 获取传感器列表
func GetProjects(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	projects, total := dao.GetProjects(pageSize, pageNum)
	code = errmessage.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmessage.Geterrmessage(code),
		"data":    projects,
		"total":   total,
	})
}
//通过项目名称获取项目
func GetProjectByNameHander(c *gin.Context){
	//获取前端传入参数
	name := c.Query("name")
	project , code := dao.GetProjectByName(name)
	c.JSON(http.StatusOK,gin.H{
		"status":  code,
		"message": errmessage.Geterrmessage(code),
		"data":    project,
	})
}

// 删除项目和他名下的传感器
func DeleteProjectAndSensors(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code = dao.DeleteProject(id)
	if code != errmessage.SUCCESS {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  code,
			"message": errmessage.Geterrmessage(code),
		})
		return
	}
	code = dao.DeleteProjectSensor(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmessage.Geterrmessage(code),
	})
}

// 暂时存在问题
func UpdateProject(c *gin.Context) {
	//id, _ := strconv.Atoi(c.Param("id"))
	var temp model.Project
	_ = c.ShouldBindJSON(&temp)
	id := temp.ID
	fmt.Println(id,id)
	fmt.Println(temp)
	data, code := dao.GetProjectByID(int(id))
	if code != errmessage.SUCCESS {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  code,
			"message": errmessage.Geterrmessage(code),
		})
		return
	}
	//_ = c.ShouldBindJSON(data)
	data = &temp
	msg, IsLegal := myValidator.MyValidate(data)
	if IsLegal != errmessage.SUCCESS {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  IsLegal,
			"message": msg,
		})
		return
	}
	fmt.Println(*data)
	dao.DeleteProject(int(id))
	dao.CreateProject(data)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmessage.Geterrmessage(code),
	})

}
