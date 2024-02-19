package alert

import (
	"fmt"
	"io/ioutil"
	"mountain/internal/dao"
	"mountain/internal/model"
	errmessage "mountain/pkg/errcode"
	"net/http"
	"strings"
)

// 利用pushplus微信公众号发送预警消息
func SendWechatMessage(Token string, alert model.Alert) {
	url := "https://www.pushplus.plus/send"
	method := "POST"
	token := Token
	title := alert.AlertName
	content := alert.AlertParameter
	template := "json"
	payload := strings.NewReader(fmt.Sprintf(`{
		"token":"%s",
		"title":"%s",
		"content":"%s",
		"template":"%s"
	}`, token, title, content, template))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("User-Agent", "Apifox/1.0.0 (https://apifox.com)")
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}

// 判断传感器数据是否异常并推送消息
func CheckSensorData(data model.SensorData, token string) {
	var alert model.Alert
	//判断数据是否设置预警
	alert, code := dao.GetAlertBySensorSN(data.SensorName)
	if code != errmessage.SUCCESS {
		return
	}
	//判断传感器数值是否异常
	if data.ValueX > alert.VXThreshold || data.ValueY > alert.VYThreshold || data.ValueH > alert.VHThreshold || data.SumX > alert.SumXThreshold || data.SumY > alert.SumYThreshold || data.SumH > alert.SumHThreshold {
		SendWechatMessage(token, alert)
		return
	} else {
		return
	}
}
