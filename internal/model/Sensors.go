package model

import (
	"gorm.io/gorm"
)

// 每个传感器序列号唯一
type Sensor struct {
	gorm.Model
	SensorSN   string  `gorm:"type:varchar(20);not null" json:"sensor_sn"` //传感器序列号
	SensorName string  `gorm:"type:varchar(20)" json:"sensor_name"`
	Longitude  float64 `gorm:"type:float;not null" json:"longitude" validate:"required,min=0,max=180" label:"经度"` //经度
	Latitude   float64 `gorm:"type:float;not null" json:"latitude" validate:"required,min=0,max=180" label:"纬度" ` //纬度
	Project    string  `gorm:"type:varchar(20)" json:"project"`
	Type       string  `gorm:"type:varchar(20)" json:"type"`
	Interval   int     `gorm:"type:int;DEFAULT:10" json:"interval" validate:"required,gte=1" label:"采集间隔"` //采集间隔
	Iswarn     string  `gorm:"type:varchar(20)" json:"iswarn"`
	Ispush     string  `gorm:"type:varchar(20)" json:"ispush"`
	Commit     string  `gorm:"type:varchar(100)" json:"commit"`
}
