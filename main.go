package main

import (
	"./config"
	"github.com/gin-gonic/gin"
	"log"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	if config.GetEnv().DEBUG {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	log.Println("env.DATABASE_IP" + config.GetEnv().SERVER_PORT)

	router := initRouter() // 初始化路由
	router.Run(":" + config.GetEnv().SERVER_PORT)
}
