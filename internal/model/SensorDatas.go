package model

import (
	"gorm.io/gorm"
	"time"
)

type SensorData struct {
	gorm.Model
	SensorName 			string  `gorm:"type:varchar(20)" json:"sensor_name"`
	CollectTime         time.Time `gorm:"type:timestamp;not null" json:"collect_time" label:"采集时间"`
	ValueX 				float64 `gorm:"type:float;not null" json:"value_x" label:"X坐标值"`
	ValueY 				float64 `gorm:"type:float;not null" json:"value_y" label:"Y坐标值"`
	ValueH 				float64 `gorm:"type:float;not null" json:"value_h" label:"H坐标值"`
	SumX				float64 `gorm:"type:float;not null" json:"sum_x"   label:"X方向累计"`
	SumY				float64 `gorm:"type:float;not null" json:"sum_y"   label:"Y方向累计"`
	SumH				float64 `gorm:"type:float;not null" json:"sum_h"   label:"H方向累计"`
	NowX				float64 `gorm:"type:float;not null" json:"now_x"   label:"X方向本次"`
	NowY				float64 `gorm:"type:float;not null" json:"now_y"   label:"Y方向本次"`
	NowH				float64 `gorm:"type:float;not null" json:"now_h"   label:"H方向本次"`

}
