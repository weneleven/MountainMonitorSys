package routes

import (
	"github.com/gin-gonic/gin"
	v1 "mountain/api/v1"
	"mountain/middleware"
	"mountain/utils"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	auth_V1 := r.Group("/api/v1")
	auth_V1.Use(middleware.JwtToken()) //需要件权的
	{
		//用户模块的路由接口
		auth_V1.PUT("user/:id", v1.EditUsers)
		auth_V1.DELETE("user/:id", v1.DeleteUser)
	}
	//公共端口
	public_V1 := r.Group("/api/v1")
	{
		public_V1.POST("user/add", v1.AddUser) //增加用户
		public_V1.GET("users", v1.GetUsers)
		public_V1.POST("login", v1.Login)
	}

	r.Run(utils.HttpPort)
}
