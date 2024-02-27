package dao

import (
	"mountain/global"
	"mountain/internal/model"
)

func GetFileInfoList(pageSize int, pageNum int) ([]model.FileInfo, int64) {
	var fileinfolist []model.FileInfo
	var total int64
	offset := (pageNum - 1) * pageSize
	if pageSize == -1 && pageNum == -1 {
		offset = -1
	}
	err := global.DBEngine.Limit(pageSize).Offset(offset).Find(&fileinfolist).Count(&total).Error
	if err != nil {
		return nil, 0
	}
	return fileinfolist, total
}
