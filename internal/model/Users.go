package model

import (
	"gorm.io/gorm"
)

// 每个用户名唯一 每个手机号唯一
type User struct {
	gorm.Model        //gorm提供的基础模型定义
	Username   string `gorm:"type:varchar(20);not null" json:"username" validate:"required,min=4,max=12" label:"用户名"`
	Password   string `gorm:"type:varchar(20);not null" json:"password" validate:"required,min=6,max=20" label:"密码"`
	Role       int    `gorm:"type:int;DEFAULT:2" json:"role" validate:"required,lte=2" label:"身份码"` //身份 gte大于等于2 lte小于等于
	Phone      string `gorm:"type:varchar(20);not null" json:"phone" validate:"required,lte=11,gte=11" label:"手机号"`
	Email      string `gorm:"type:varchar(30)" json:"email" `
	Department string `gorm:"type:varchar(20)" json:"department" `
	Sex        int    `gorm:"type:int;DEFAULT:1" json:"sex" validate:"required,lte=2,gte=1" label:"性别码"` //1为男性 2为女性
}


