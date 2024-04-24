package model

import (
	"gorm.io/gorm"
)

type Displacement struct {
	gorm.Model
	SensorName string  `gorm:"type:varchar(20)" json:"sensor_name"`
	Date       string  `gorm:"type:varchar(20);not null" json:"date"`
	Nd         float64 `gorm:"type:float;not null" json:"nd"`
	Ed         float64 `gorm:"type:float;not null" json:"ed"`
	Hd         float64 `gorm:"type:float;not null" json:"hd"`
	Nvel       float64 `gorm:"type:float;not null" json:"nvel"`
	Evel       float64 `gorm:"type:float;not null" json:"evel"`
	Hvel       float64 `gorm:"type:float;not null" json:"hvel"`
}
