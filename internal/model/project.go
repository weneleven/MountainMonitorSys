package model

import (
	"gorm.io/gorm"
)

type Project struct {
	gorm.Model                  // GORM基础模型
	ProjectName          string `gorm:"type:varchar(100);not null" json:"project_name" validate:"required" label:"项目名称"`                  //项目名称
	ProjectShortName     string `gorm:"type:varchar(50);not null" json:"project_short_name" validate:"required,max=1000" label:"项目简称"`    //项目简称 最多50字
	ProjectType          string `gorm:"type:varchar(50);not null" json:"project_type" validate:"required" label:"项目类型"`                   //项目类型
	Location             string `gorm:"type:varchar(100);not null" json:"location" validate:"required" label:"项目位置"`                      //项目位置
	AlertsEnabled        bool   `gorm:"type:boolean;DEFAULT:false" json:"alerts_enabled" label:"预警启用状态"`                                  //预警启用 true / false
	AutoAlert            bool   `gorm:"type:boolean;DEFAULT:false" json:"auto_alert" label:"自动推送预警"`                                      //推自动送预警 true / false
	OfflinePushEnabled   bool   `gorm:"type:boolean;DEFAULT:false" json:"offline_push_enabled" label:"离线推送"`                              //离线推送是否启用
	OfflinePushFrequency int    `gorm:"type:int;DEFAULT:86400" json:"offline_push_frequency" validate:"required,gte=0" label:"离线推送频率(s)"` //离散推送频率,单位秒
	MapScale             int    `gorm:"type:int" json:"map_scale" validate:"required,gte=0" label:"地图比例"`                                 //地图比例
	SmsSignature         string `gorm:"type:varchar(100)" json:"sms_signature" label:"短信签名"`                                              //短信签名
	Remark               string `gorm:"type:text" json:"remark" label:"备注信息"`                                                             //备注
	ProjectDescription   string `gorm:"type:text" json:"project_description" validate:"max=1000" label:"项目简介"`                            //项目简介 最多1000字
	ObserverName         string `gorm:"type:varchar(50)" json:"observer_name" label:"观察者姓名"`                                              //观察者
	ValidatorName        string `gorm:"type:varchar(50)" json:"validator_name" label:"校验者姓名"`                                             //校验者
	CalculatorName       string `gorm:"type:varchar(50)" json:"calculator_name" label:"计算者姓名"`                                            //计算者
	ReviewerName         string `gorm:"type:varchar(50)" json:"reviewer_name" label:"审核者姓名"`                                              //审核者
	InspectionUnitName   string `gorm:"type:varchar(100)" json:"inspection_unit_name" label:"检测单位名称"`                                     //单位
	UserID               uint   `gorm:"not null" json:"user_id" validate:"required" label:"用户ID"`                                         //所属用户ID
}
type Project1 struct {
	ProjectID            int    //项目ID
	Name                 string //项目名称
	Type                 string //项目类型
	Address              string //项目地址
	AlertsEnabled        bool   //预警启用状态
	AlertMethod          string //推送预警方式
	OfflinePushEnabled   bool   //离线推送
	OfflinePushFrequency int    //离线推送频率
	MapScale             int    //地图比例
	ReviewerName         string //审核者
	UserID               uint   //所属用户ID
}
