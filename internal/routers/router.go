package routes

import (
	"github.com/gin-gonic/gin"
	v1 "mountain/api/v1"
	"mountain/global"
	"mountain/internal/middleware"
	"mountain/pkg/logger"
)

func InitRouter() {
	gin.SetMode(global.ServerSetting.AppMode)
	r := gin.Default()
	r.Use(logger.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())
	auth_V1 := r.Group("/api/v1")
	auth_V1.Use(middleware.JwtToken()) //需要鉴权的
	{
		//用户模块的路由接口
		auth_V1.PUT("user/:id", v1.EditUsers)
		auth_V1.DELETE("user/:id", v1.DeleteUser)

		//项目模块的路由接口
		auth_V1.POST("project/add",v1.AddProject)
		auth_V1.GET("projects", v1.GetProjects)
		auth_V1.DELETE("project/:id", v1.DeleteProjectAndSensors)
		auth_V1.PUT("project/:id", v1.UpdateProject)

		//传感器模块的路由接口
		auth_V1.POST("sensor/add", v1.AddSensor)
		auth_V1.GET("sensors", v1.GetSensors)
		auth_V1.DELETE("sensor/:id", v1.DeleteSensor)
		auth_V1.PUT("sensor/:id", v1.EditSensor)
	}
	//公共端口
	public_V1 := r.Group("/api/v1")
	{
		public_V1.POST("user/add", v1.AddUser) //增加用户
		public_V1.GET("users", v1.GetUsers)
		public_V1.POST("login", v1.Login)
	}

	r.Run(global.ServerSetting.HttpPort)

}
