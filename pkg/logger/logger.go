package logger

import (
	"fmt"
	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"math"
	"os"
	"time"
)

func Logger() gin.HandlerFunc {
	filePath := "log/"
	src, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println("err:", err)
	}
	logger := logrus.New()
	logger.Out = src
	logger.SetLevel(logrus.DebugLevel)
	logWriter, _ := rotatelogs.New(
		filePath+"%Y%m%d.log",
		rotatelogs.WithMaxAge(7*time.Hour),
		rotatelogs.WithRotationTime(time.Hour),
		rotatelogs.WithClock(rotatelogs.Local),
	)
	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.DebugLevel: logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.FatalLevel: logWriter,
		logrus.PanicLevel: logWriter,
		logrus.TraceLevel: logWriter,
		logrus.WarnLevel:  logWriter,
	}
	Hook := lfshook.NewHook(writeMap, &logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	logger.AddHook(Hook)

	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		stopTime := time.Since(startTime)
		spendTime := fmt.Sprintf("%d ms", int(math.Ceil(float64((stopTime.Nanoseconds())/1000000.0))))
		hostName, err := os.Hostname()
		if err != nil {
			hostName = "unknow"
		}
		statusCode := c.Writer.Status()    //状态码
		clienIp := c.ClientIP()            //客户端
		userAgent := c.Request.UserAgent() //用户代理
		dataSize := c.Writer.Size()        //数据大小
		if dataSize < 0 {
			dataSize = 0
		}
		method := c.Request.Method   //方法
		path := c.Request.RequestURI //请求路径

		entry := logger.WithFields(logrus.Fields{
			"Hostname":  hostName,
			"status":    statusCode,
			"spendtime": spendTime,
			"Ip":        clienIp,
			"Agent":     userAgent,
			"datasize":  dataSize,
			"path":      path,
			"Method":    method,
		})
		if len(c.Errors) > 0 {
			entry.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
		}
		if statusCode >= 500 {
			entry.Error()
		} else if statusCode >= 400 {
			entry.Warn()
		} else {
			entry.Info()
		}
	}
}
