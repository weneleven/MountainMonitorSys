package dao

import (
	"mountain/global"
	"mountain/internal/model"
	errmessage "mountain/pkg/errcode"
)

// 新增传感器
func CreatSensor(data *model.Sensor) int {
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
func GetSensorByID(id int) (*model.Sensor, int) {
	var sensor model.Sensor
	result := global.DBEngine.Where("id = ?", id).First(&sensor)
	if result.Error != nil {
		return nil, errmessage.ERROR_SENSOR_NOT_EXIST
	}
	return &sensor, errmessage.SUCCESS
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

// 更新传感器
func UpdateSensor(id int, data *model.Sensor) int {
	var sensor model.Sensor
	maps := make(map[string]interface{}) //设备号无法更新
	if sensor.Longitude != 0 {
		maps["longitude"] = sensor.Longitude
	}
	if sensor.Latitude != 0 {
		maps["latitude"] = sensor.Latitude
	}
	if sensor.Altitude != 0 {
		maps["altitude"] = sensor.Altitude
	}
	if sensor.Address != "" {
		maps["address"] = sensor.Address
	}
	if sensor.AcquisitionInterval != 0 {
		maps["acquisition_interval"] = sensor.AcquisitionInterval
	}
	if sensor.ArchiveStatus != 0 {
		maps["archive_status"] = sensor.ArchiveStatus
	}
	if sensor.IsWarning != 0 {
		maps["is_warning"] = sensor.IsWarning
	}
	if sensor.WarningInterval != 0 {
		maps["warning_interval"] = sensor.WarningInterval
	}
	if sensor.Commit != "" {
		maps["commit"] = sensor.Commit
	}
	result := global.DBEngine.Model(&sensor).Where("id = ?", id).Updates(data)
	if result.Error != nil {
		return errmessage.ERROR
	}
	return errmessage.SUCCESS
}
