package websocket

import (
	"fmt"
	"mountain/internal/alert"
	"mountain/internal/dao"
	"mountain/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func SensorHandler(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	defer ws.Close()

	for {
		messageType, p, err := ws.ReadMessage()
		if err != nil {
			fmt.Println("***error***")
			fmt.Println(err)
			return
		}
		fmt.Println(messageType)
		fmt.Printf("***Received message: %s\n", p)
		var data model.SensorData
		err = ws.ReadJSON(&data)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println("***data:***")
		fmt.Println(data)
		//将数据存入数据库
		dao.AddSensorData(&data)
		//检测数据预警
		alert.CheckSensorData(data, "d4e5428286fe42058c0a03ff450cb790") //这里的token是我自己的token，需要自己申请 具体到每个用户
	}
}
