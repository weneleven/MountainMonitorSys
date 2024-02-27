package model

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"mountain/pkg/setting"
	"time"
)

func NewDBEngine(databaseSetting *setting.DatabaseSettingS) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		databaseSetting.UserName,
		databaseSetting.Password,
		databaseSetting.Host,
		databaseSetting.DBName,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("连接数据库失败，请检查参数:", err)
	}
	sqlDB, err := db.DB()

	if err != nil {
		fmt.Println("连接数据库失败，请检查参数:", err)
	}
	// 自动迁移用户的字段
	err = db.AutoMigrate(&User{})
	err = db.AutoMigrate(&Project{})
	err = db.AutoMigrate(&Sensor{})
	err = db.AutoMigrate(&SensorData{})
	err = db.AutoMigrate(&FileInfo{})
	err = db.AutoMigrate(&Alert{})
	if err != nil {
		return nil, fmt.Errorf("自动迁移失败: %w", err)
	}
	//...
	sqlDB.SetMaxIdleConns(10)                  //连接池中最大闲置连接池数 (没有请求的时候连接池需要的最大连接数量)
	sqlDB.SetMaxOpenConns(100)                 //数据库最大连接数量
	sqlDB.SetConnMaxLifetime(10 * time.Second) //链接最大可复用时间

	return db, nil
}

// 添加TDengine连接
func NewDb() (*sql.DB, error) {
	var taosDSN = "root:taosdata@tcp(124.70.83.36:6030)/mountain"
	taos, err := sql.Open("taosSql", taosDSN)
	if err != nil {
		fmt.Println("failed to connect TDengine, err:", err)
	}
	fmt.Println("connected")
	return taos, nil
}
