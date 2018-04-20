package logger

import (
	"os"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"morningo/config"
)

func init()  {
	if config.GetEnv().ACCESS_LOG {
		file, err := os.OpenFile(config.GetEnv().ACCESS_LOG_PATH, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalln(err)
		}
		if config.GetEnv().DEBUG {
			gin.DefaultWriter = io.MultiWriter(file, os.Stdout)
		} else {
			gin.DefaultWriter = io.MultiWriter(file)
		}
	}
}
