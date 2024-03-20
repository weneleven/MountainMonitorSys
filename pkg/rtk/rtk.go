package rtk

import (
	"errors"
	"mountain/internal/model"
)

// GNSS数据解析函数
func ParseGNSSData(data model.SensorData) (model.SensorData, error) {
	// 检查传入的数据是否有效
	if data.SensorName == "" {
		return model.SensorData{}, errors.New("invalid data: missing sensor name")
	}

	// 这里添加你的GNSS数据解析逻辑
	// 例如，你可能需要解析data.ValueX, data.ValueY, data.ValueH等字段
	// 解析后的结果存储在一个新的model.SensorData对象中

	parsedData, err := rtkPositioning(data)
	if err != nil {
		return model.SensorData{}, errors.New("Fail")
	}
	return parsedData, nil
}

// RTK定位函数
func rtkPositioning(data model.SensorData) (model.SensorData, error) {
	// 检查传入的数据是否有效
	if data.SensorName == "" {
		return model.SensorData{}, errors.New("invalid data: missing sensor name")
	}

	ans := model.SensorData{
		Model:       data.Model,
		SensorName:  data.SensorName,
		CollectTime: data.CollectTime,
		ValueX:      data.ValueX,
		ValueY:      data.ValueY,
		ValueH:      data.ValueH,
		SumX:        data.SumX,
		SumY:        data.SumY,
		SumH:        data.SumH,
		NowX:        data.NowX,
		NowY:        data.NowY,
		NowH:        data.NowH,
	}

	return ans, nil
}
