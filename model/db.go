package model

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"mountain/utils"
	"time"
)

var db *gorm.DB
var err error

func InitDb() {
	db, err = gorm.Open(utils.Db, fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		utils.DbUser,
		utils.DbPassWord,
		utils.DbHost,
		utils.DbPort,
		utils.DbName,
	))
	if err != nil {
		fmt.Println("连接数据库失败，请检查参数:", err)
	}
	//自动迁移用户的字段
	db.AutoMigrate(&User{})

	db.DB().SetMaxIdleConns(10)                  //连接池中最大闲置连接池数 (没有请求的时候连接池需要的最大连接数量)
	db.DB().SetMaxOpenConns(100)                 //数据库最大连接数量
	db.DB().SetConnMaxLifetime(10 * time.Second) //链接最大可复用时间

	//db.Close()
}
