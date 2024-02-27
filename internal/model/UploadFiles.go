package model

import "gorm.io/gorm"

type FileInfo struct {
	gorm.Model
	FileName          string `gorm:"type:varchar(100);not null" json:"file_name" label:"文件名称"`
	URL     string `gorm:"type:varchar(100);not null" json:"url"  label:"url"`
	Extension     string `gorm:"type:varchar(100);not null" json:"extension"  label:"文件后缀"`
}
