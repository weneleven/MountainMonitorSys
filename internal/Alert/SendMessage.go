package Alert

import (
	"fmt"
	"io/ioutil"
	"mountain/internal/model"
	"net/http"
	"strings"
)

// 利用pushplus微信公众号发送预警消息
func SendWechatMessage(user model.User) {

	url := "https://www.pushplus.plus/send"
	method := "POST"

	payload := strings.NewReader(`{
    "token":"be8652e98bd04139a13e1bffcd6d0f71",
    "title":"标题",
    "content":"{'name':'名称','size':'大小','number':'数量'}",
    "template":"json"
}`)

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
