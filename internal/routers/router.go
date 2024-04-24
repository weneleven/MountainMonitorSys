package routes

import (
	v1 "mountain/api/v1"
	"mountain/global"
	"mountain/internal/middleware"
	"mountain/internal/websocket"
	"mountain/pkg/logger"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(global.ServerSetting.AppMode)
	r := gin.Default()
	r.Use(logger.Logger())
	r.Use(gin.Recovery())
	// 使用 CORS 中间件配置
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"} // 允许的前端域，可以是具体的域名，也可以是通配符 "*"
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Content-Length", "Accept-Encoding", "Authorization"}
	config.AllowCredentials = true
	config.MaxAge = 12 * time.Hour
	r.Use(cors.New(config)) // 在所有路由之前使用 CORS 中间件
	auth_V1 := r.Group("/api/v1")
	auth_V1.Use(middleware.JwtToken()) //需要鉴权的
	{
		//用户模块
		auth_V1.POST("user/add", v1.AddUser) //增加用户
		auth_V1.PUT("user/:id", v1.EditUsers)
		auth_V1.DELETE("user/:id", v1.DeleteUser)

		//项目模块
		auth_V1.POST("project/add", v1.AddProject)
		auth_V1.GET("projects", v1.GetProjects)
		auth_V1.DELETE("project/:id", v1.DeleteProjectAndSensors)
		auth_V1.PUT("project/updateProject", v1.UpdateProject)
		auth_V1.GET("project/getByName", v1.GetProjectByNameHander)
		auth_V1.DELETE("project/deleteById", v1.DeleteProjects)
		//传感器模块
		auth_V1.POST("sensor/add", v1.AddSensor)
		auth_V1.GET("sensors", v1.GetSensors)
		auth_V1.DELETE("sensor/:id", v1.DeleteSensor)
		auth_V1.PUT("sensor/:id", v1.EditSensor)
		//数据处理与展示模块
		auth_V1.GET("data/get", v1.GetSensorData)
		//auth_V1.POST("data/add", v1.AddSensorData)
		auth_V1.GET("data/getBySN", v1.GetSensorDataBySN)
		auth_V1.POST("data/upload", v1.UploadFile)
		auth_V1.GET("data/getfileinfo", v1.GetFileInfo)
		//预警模块
		auth_V1.GET("alert/get", v1.GetAlerts)
		auth_V1.POST("alert/add", v1.AddAlert)
		auth_V1.DELETE("alert/:id", v1.DeleteAlert)
		auth_V1.PUT("alert/:id", v1.EditAlert)
		auth_V1.GET("warn/get", v1.GetWarn)
		auth_V1.GET("warn/getBySn", v1.GetWarnBySN)
		auth_V1.GET("displacement/get", v1.GetDisplace)
		auth_V1.GET("displacement/getBySn", v1.GetDisplaceBySN)
	}
	//公共端口
	public_V1 := r.Group("/api/v1")
	{
		public_V1.GET("users", v1.GetUsers)
		public_V1.POST("login", v1.Login)

	}

	r.GET("/ws/sensor", websocket.SensorHandler)

	r.Run(global.ServerSetting.HttpPort)

}
