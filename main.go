package main

import (
	"mountain/model"
	"mountain/routes"
)

func main() { //我能在这看到是我写的
	//引用数据库
	model.InitDb()
	//初始化路由
	routes.InitRouter()
}
