package main

import (
	"morningo/config"
	"github.com/gin-gonic/gin"
	"runtime"
	"os"
	"log"
	"io"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	if config.GetEnv().DEBUG {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	if config.GetEnv().ACCESS_LOG {
		file, err := os.OpenFile("storage/logs/access.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalln(err)
		}
		if config.GetEnv().DEBUG {
			gin.DefaultWriter = io.MultiWriter(file, os.Stdout)
		} else {
			gin.DefaultWriter = io.MultiWriter(file)
		}
	}

	router := initRouter() // 初始化路由
	router.Run(":" + config.GetEnv().SERVER_PORT)
}
