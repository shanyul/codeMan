package server

import (
	"chujian-api/internal/model"
	"chujian-api/pkg/gredis"
	"chujian-api/pkg/logging"
	"chujian-api/pkg/setting"
	"chujian-api/pkg/util"
	"chujian-api/server/router"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func StartApp() {
	// 初始化配置
	setting.Setup()
	// 数据库连接
	//model.SetUp()
	// 初始化日志
	logging.Setup()
	// 初始化 redis
	//gredis.Setup()
	// 初始化配置
	util.Setup()
	// 运行服务
	startServer()
}

func ExitApp() {
	model.CloseDB()
	gredis.Close()
}

func startServer() {
	// 设置运行模式
	gin.SetMode(setting.ServerSetting.RunMode)

	routersInit := router.InitRouter()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] start http server listening %s", endPoint)
	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v", err)
	}
}
