package websocket

import (
	"fmt"
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
		//dao.AddSensorData(&data) //将数据存入数据库
	}
}
