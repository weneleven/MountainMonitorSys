package dao

import (
	"mountain/global"
	"mountain/internal/model"
	errmessage "mountain/pkg/errcode"
)

// 新增预警信息
func AddAlert(data *model.Alert) int {
	result := global.DBEngine.Create(data)
	if result.Error != nil {
		return errmessage.ERROR
	}
	return errmessage.SUCCESS
}

// 删除预警信息
func DeleteAlert(id int) int {
	var alert model.Alert
	result := global.DBEngine.Where("id = ?", id).Delete(&alert)
	if result.Error != nil {
		return errmessage.ERROR
	}
	return errmessage.SUCCESS
}

// 获取预警信息列表
func GetAlerts(pageSize int, pageNum int) ([]model.Alert, int64) {
	var alerts []model.Alert
	var total int64
	offset := (pageNum - 1) * pageSize
	if pageSize == -1 && pageNum == -1 {
		offset = -1
	}
	err := global.DBEngine.Limit(pageSize).Offset(offset).Find(&alerts).Count(&total).Error
	if err != nil {
		return nil, 0
	}
	return alerts, total
}

// 编辑预警信息
func EditAlert(id int, data *model.Alert) int {
	result := global.DBEngine.Model(&model.Alert{}).Where("id = ?", id).Updates(data)
	if result.Error != nil {
		return errmessage.ERROR
	}
	return errmessage.SUCCESS
}
