package main

import (
	"github.com/gin-gonic/gin"
	_ "morningo/modules/log" // 日志
	// _ "morningo/modules/schedule" // 定时任务
	"runtime"
	"morningo/config"
	"morningo/modules/server"
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
