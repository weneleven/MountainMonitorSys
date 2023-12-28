package main

import (
	"fmt"
	_ "github.com/taosdata/driver-go/v3/taosRestful"
	_ "github.com/taosdata/driver-go/v3/taosSql"
	"log"
	"mountain/global"
	"mountain/internal/model"
	routes "mountain/internal/routers"
	"mountain/pkg/setting"
	"time"
)

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
	err = setupDBEngine()
	//err = setupDb()
	if err != nil {
		fmt.Println(err)
		log.Fatalf("init.setupDBEngine err: %v", err)
	}
	//dao.Readfile()
}

func main() {
	//初始化路由
	routes.InitRouter()
}

func setupSetting() error {
	setting, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}

	global.ServerSetting.ReadTimeout = time.Second
	global.ServerSetting.WriteTimeout = time.Second
	return nil
}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}

	return nil
}
func setupDb() error {
	var err error
	global.Db, err = model.NewDb()
	if err != nil {
		return err
	}

	return nil
}
