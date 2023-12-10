package model

import (
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
	if err != nil {
		return nil, fmt.Errorf("自动迁移失败: %w", err)
	}
	//...
	sqlDB.SetMaxIdleConns(10)                  //连接池中最大闲置连接池数 (没有请求的时候连接池需要的最大连接数量)
	sqlDB.SetMaxOpenConns(100)                 //数据库最大连接数量
	sqlDB.SetConnMaxLifetime(10 * time.Second) //链接最大可复用时间

	return db, nil
}
