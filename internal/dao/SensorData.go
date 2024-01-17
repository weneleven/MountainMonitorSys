package dao

import (
	"encoding/csv"
	"fmt"
	"mountain/global"
	"mountain/internal/model"
	"os"
	"strconv"
	"time"
)

func GetSensorData() ([]model.SensorData, int64) {
	var datas []model.SensorData
	var total int64
	err := global.DBEngine.Find(&datas).Count(&total).Error
	if err != nil {
		return nil, 0
	}
	return datas, total
}
//从文件中批量读取采集数据存入mysql
func Readfile() {
	// 读取 CSV 文件
	file, err := os.Open("/mycode/data/data.csv")
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
		data := model.SensorData{
			SensorName:  record[0],
			CollectTime: parseTime(record[1]),
			ValueX:      parseFloat(record[2]),
			ValueY:      parseFloat(record[3]),
			ValueH:      parseFloat(record[4]),
			SumX:        parseFloat(record[5]),
			SumY:        parseFloat(record[6]),
			SumH:        parseFloat(record[7]),
			NowX:        parseFloat(record[8]),
			NowY:        parseFloat(record[9]),
			NowH:        parseFloat(record[10]),
		}

		// 插入数据
		result := global.DBEngine.Create(&data)
		if result.Error != nil {
			fmt.Println("插入数据失败:", result.Error)
		}
	}

	fmt.Println("数据插入成功")
}
func parseTime(timeStr string) time.Time {
	t, _ := time.Parse("2006-01-02 15:04:05", timeStr)
	return t
}

func parseFloat(floatStr string) float64 {
	// 解析浮点数字符串
	f, _ := strconv.ParseFloat(floatStr, 64)
	return f
}



