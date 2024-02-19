package model

import (
	"gorm.io/gorm"
)

type Alert struct {
	gorm.Model
	SensorSN       string  `gorm:"type:varchar(20)" json:"sensor_sn" validate:"required,max=20" label:"报警传感器"`
	AlertName      string  `gorm:"type:varchar(50)" json:"alert_name" validate:"required,max=50" label:"报警名称"`
	AlertParameter string  `gorm:"type:varchar(50)" json:"alert_parameter" validate:"required,max=50" label:"报警字段"`
	VXThreshold    float64 `json:"vx_threshold" validate:"required" label:"X值阈值"`
	VYThreshold    float64 `json:"vy_threshold" validate:"required" label:"Y值阈值"`
	VHThreshold    float64 `json:"vh_threshold" validate:"required" label:"H值阈值"`
	SumXThreshold  float64 `json:"sum_x_threshold" validate:"required" label:"X方向累计阈值"`
	SumYThreshold  float64 `json:"sum_y_threshold" validate:"required" label:"Y方向累计阈值"`
	SumHThreshold  float64 `json:"sum_h_threshold" validate:"required" label:"H方向累计阈值"`
	Frequency      int     `json:"frequency" validate:"required,min=1" label:"报警次数"`
}
