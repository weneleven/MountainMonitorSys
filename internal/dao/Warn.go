package dao

import (
	"encoding/csv"
	"fmt"
	"mountain/global"
	"mountain/internal/model"
	errmessage "mountain/pkg/errcode"
	"os"
)

// 新增传感器数据
func AddWarn(data *model.Warn) int {
	result := global.DBEngine.Create(data)
	if result.Error != nil {
		return errmessage.ERROR
	}
	return errmessage.SUCCESS
}

// 通过设备号获取采集数据
func GetWarnBySensorSN(sn string) ([]model.Warn, int64, int) {
	var datas []model.Warn
	var total int64
	result := global.DBEngine.Where("sensor_name = ?", sn).Find(&datas).Count(&total)
	if result.Error != nil {
		return nil, 0, errmessage.ERROR_SENSOR_NOT_EXIST
	}
	if datas == nil {
		return nil, 0, errmessage.ERROR_SENSOR_NOT_EXIST
	}
	return datas, total, errmessage.SUCCESS
}

// 获取采集数据
func GetWarns() ([]model.Warn, int64) {
	var datas []model.Warn
	var total int64
	err := global.DBEngine.Find(&datas).Count(&total).Error
	if err != nil {
		return nil, 0
	}
	return datas, total
}
func ReadWarn() {
	// 读取 CSV 文件
	file, err := os.Open("/code/mountain-sys/dataSent/warns.csv")
	if err != nil {
		fmt.Println("打开文件失败:", err)
		return
	}
	// 解析 CSV 文件
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("解析CSV文件失败:", err)
		return
	}
	for _, record := range records {
		data := model.Warn{
			SensorName:         record[0],
			Today:              record[1],
			Yesterday:          record[2],
			TodayCombinexy:     ParseFloat(record[3]),
			YesterdayCombinexy: ParseFloat(record[4]),
			WeekAvgVxy:         ParseFloat(record[5]),
			TodayCombineh:      ParseFloat(record[6]),
			TodayVh:            ParseFloat(record[7]),
			YesterdayCombineh:  ParseFloat(record[8]),
			YesterdayVh:        ParseFloat(record[9]),
			Xya:                ParseFloat(record[10]),
			Ha:                 ParseFloat(record[11]),
			XyWarnInfo:         record[12],
			XyFlag:             ParseFloat(record[13]),
			HWarnInfo:          record[14],
			HFlag:              ParseFloat(record[15]),
		}

		// 插入数据
		result := global.DBEngine.Create(&data)
		if result.Error != nil {
			fmt.Println("插入数据失败:", result.Error)
		}
	}
	fmt.Println("数据插入成功")
}
