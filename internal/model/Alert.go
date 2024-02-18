package model

import "google.golang.org/protobuf/types/known/timestamppb"

type Alert struct {
	SensorSN       string                //报警传感器
	AlertName      string                //报警名称
	AlertParameter string                //报警字段
	AlertValue     float64               //报警时的值
	Threshold      float64               //阈值
	Ts             timestamppb.Timestamp //时间戳
	status         string                //处理状态
	frequency      int                   //报警次数
}
