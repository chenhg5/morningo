package main

import (
	"github.com/gin-gonic/gin"
	_ "morningo/module/logger" // 日志
	// _ "morningo/module/schedule" // 定时任务
	"runtime"
	"morningo/config"
	"morningo/module/server"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	if config.GetEnv().DEBUG {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	router := initRouter()

	server.Run(router)
}
