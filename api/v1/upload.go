package v1

import (
	"github.com/gin-gonic/gin"
	"io"
	"mountain/global"
	"mountain/internal/dao"
	"mountain/internal/model"
	errmessage "mountain/pkg/errcode"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

func UploadFile(c *gin.Context)  {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无法获取文件"})
		return
	}
	defer file.Close()
	// 创建目标文件
	dst, err := os.Create("/mycode/mountain-sys/dataSent/" + header.Filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法创建文件"})
		return
	}
	defer dst.Close()
	// 将上传的文件内容拷贝到目标文件中
	_, err = io.Copy(dst, file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法写入文件内容"})
		return
	}
	// 获取文件后缀
	ext := filepath.Ext(header.Filename)
	// 创建文件信息记录
	fileInfo := model.FileInfo{
		FileName:  header.Filename,
		URL:       "/uploaded_files/" + header.Filename,
		Extension: ext,
	}
	// 将文件信息存入数据库
	result := global.DBEngine.Create(&fileInfo)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法存储文件信息"})
		return
	}
	// 返回上传成功的消息和文件URL
	c.JSON(http.StatusOK, gin.H{"message": "文件上传成功", "file_name": header.Filename, "url": "/mycode/mountain-sys/dataSent/" + header.Filename,"extension":ext})
}

func GetFileInfo(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	fileinfolist, total := dao.GetFileInfoList(pageSize, pageNum)
	code = errmessage.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmessage.Geterrmessage(code),
		"data":    fileinfolist,
		"total":   total,
	})
}
