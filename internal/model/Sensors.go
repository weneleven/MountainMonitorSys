package model

import "gorm.io/gorm"

// 每个传感器序列号唯一
type Sensor struct {
	gorm.Model
	SensorSN            string  `gorm:"type:varchar(20);not null" json:"sensor_sn"`                                             //传感器序列号
	Longitude           float64 `gorm:"type:float;not null" json:"longitude" validate:"required,min=0,max=180" label:"经度"`      //经度
	Latitude            float64 `gorm:"type:float;not null" json:"latitude" validate:"required,min=0,max=180" label:"纬度" `      //纬度
	Altitude            float64 `gorm:"type:float;not null" json:"altitude"`                                                    //海拔
	Address             string  `gorm:"type:varchar(20)" json:"address"`                                                        //地址
	AcquisitionInterval int     `gorm:"type:int;DEFAULT:10" json:"acquisition_interval" validate:"required,gte=1" label:"采集间隔"` //采集间隔
	ArchiveStatus       int     `gorm:"type:int;DEFAULT:1" json:"archive_status" validate:"required,lte=1,gte=0" label:"归档状态"`  //0为未归档 1为已归档
	IsWarning           int     `gorm:"type:int;DEFAULT:1" json:"is_warning" validate:"required,lte=1,gte=0" label:"是否报警"`      //0为未报警 1为已报警
	WarningInterval     int     `gorm:"type:int;DEFAULT:10" json:"warning_interval" validate:"required,gte=1" label:"报警间隔"`     //报警间隔
	Commit              string  `gorm:"type:varchar(20)" json:"commit" validate:"required,min=0,max=100" label:"备注"`            //备注
	ProjectID           int     `gorm:"type:int;not null" json:"project_id" validate:"required,gte=1" label:"项目ID"`             //项目ID
}
