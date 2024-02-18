package dao

import (
	"mountain/global"
	"mountain/internal/model"
	errmessage "mountain/pkg/errcode"
)

// 检查是否存在此传感器
func CheckSensorSN(sn string) int {
	var sensor model.Sensor
	global.DBEngine.Select("id").Where("sensor_sn = ?", sn).First(&sensor)
	if sensor.ID > 0 {
		return errmessage.ERROR_SENSORSN_USED
	}
	return errmessage.SUCCESS
}

// 新增传感器
func AddSensor(data *model.Sensor) int {
	err := global.DBEngine.Create(data).Error
	if err != nil {
		return errmessage.ERROR
	}
	return errmessage.SUCCESS
}

// 获取传感器列表
func GetSensors(pageSize int, pageNum int) ([]model.Sensor, int64) {
	var sensors []model.Sensor
	var total int64
	offset := (pageNum - 1) * pageSize
	if pageSize == -1 && pageNum == -1 {
		offset = -1
	}
	err := global.DBEngine.Limit(pageSize).Offset(offset).Find(&sensors).Count(&total).Error
	if err != nil {
		return nil, 0
	}
	return sensors, total
}

// 返回此ID的传感器
func GetSensorByID(id int) (model.Sensor, int) {
	var sensor model.Sensor
	err := global.DBEngine.Where("id = ?", id).First(&sensor).Error
	if err != nil {
		return sensor, errmessage.ERROR_SENSOR_NOT_EXIST
	}
	return sensor, errmessage.SUCCESS
}

// 删除传感器
func DeleteSensor(id int) int {
	var sensor model.Sensor
	result := global.DBEngine.Where("id = ?", id).Delete(&sensor)
	if result.Error != nil {
		return errmessage.ERROR
	}
	return errmessage.SUCCESS
}

// 编辑传感器
func EditSensor(id int, data *model.Sensor) int {
	//获取数据库中的传感器
	err := global.DBEngine.Model(&model.Sensor{}).Where("id = ?", id).Updates(data).Error
	if err != nil {
		return errmessage.ERROR
	}
	return errmessage.SUCCESS
}
