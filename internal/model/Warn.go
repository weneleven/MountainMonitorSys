package model

import (
	"gorm.io/gorm"
)

type Warn struct {
	gorm.Model
	SensorName         string  `gorm:"type:varchar(20)" json:"sensor_name"`
	Today              string  `gorm:"type:varchar(20);not null" json:"today"`
	Yesterday          string  `gorm:"type:varchar(20);not null" json:"yesterday"`
	TodayCombinexy     float64 `gorm:"type:float;not null" json:"today_combinexy"`
	YesterdayCombinexy float64 `gorm:"type:float;not null" json:"yesterday_combinexy"`
	WeekAvgVxy         float64 `gorm:"type:float;not null" json:"week_avg_vxy"`
	TodayCombineh      float64 `gorm:"type:float;not null" json:"today_combineh"`
	TodayVh            float64 `gorm:"type:float;not null" json:"today_vh"`
	YesterdayCombineh  float64 `gorm:"type:float;not null" json:"yesterday_combineh"`
	YesterdayVh        float64 `gorm:"type:float;not null" json:"yesterday_vh"`
	Xya                float64 `gorm:"type:float;not null" json:"xya"`
	Ha                 float64 `gorm:"type:float;not null" json:"ha"`
	XyWarnInfo         string  `gorm:"type:varchar(20)" json:"xy_warn_info"`
	XyFlag             float64 `gorm:"type:float;not null" json:"xy_flag"`
	HWarnInfo          string  `gorm:"type:varchar(20)" json:"h_warn_info"`
	HFlag              float64 `gorm:"type:float;not null" json:"h_flag"`
}
