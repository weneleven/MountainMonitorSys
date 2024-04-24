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
func AddDisplace(data *model.Displacement) int {
	result := global.DBEngine.Create(data)
	if result.Error != nil {
		return errmessage.ERROR
	}
	return errmessage.SUCCESS
}

// 通过设备号获取采集数据
func GetDisplaceBySensorSN(sn string) ([]model.Displacement, int64, int) {
	var datas []model.Displacement
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
func GetDisplace() ([]model.Displacement, int64) {
	var datas []model.Displacement
	var total int64
	err := global.DBEngine.Find(&datas).Count(&total).Error
	if err != nil {
		return nil, 0
	}
	return datas, total
}
func ReadDisplace() {
	// 读取 CSV 文件
	file, err := os.Open("/code/mountain-sys/dataSent/displacement.csv")
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
		data := model.Displacement{
			SensorName: record[0],
			Date:       record[1],
			Nd:         ParseFloat(record[2]),
			Ed:         ParseFloat(record[3]),
			Hd:         ParseFloat(record[4]),
			Nvel:       ParseFloat(record[5]),
			Evel:       ParseFloat(record[6]),
			Hvel:       ParseFloat(record[7]),
		}

		// 插入数据
		result := global.DBEngine.Create(&data)
		if result.Error != nil {
			fmt.Println("插入数据失败:", result.Error)
		}
	}
	fmt.Println("数据插入成功")
}
