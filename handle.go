package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"io"
	"runtime/debug"
	"time"
	"gin-template/config"
	"os"
)

var (
	//green        = string([]byte{27, 91, 57, 55, 59, 52, 50, 109})
	//white        = string([]byte{27, 91, 57, 48, 59, 52, 55, 109})
	//yellow       = string([]byte{27, 91, 57, 55, 59, 52, 51, 109})
	//red          = string([]byte{27, 91, 57, 55, 59, 52, 49, 109})
	//blue         = string([]byte{27, 91, 57, 55, 59, 52, 52, 109})
	//magenta      = string([]byte{27, 91, 57, 55, 59, 52, 53, 109})
	//cyan         = string([]byte{27, 91, 57, 55, 59, 52, 54, 109})
	//reset        = string([]byte{27, 91, 48, 109})
	//disableColor = false
	DefaultWriter io.Writer
)

func handleErrors() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {

				file, openErr := os.OpenFile("storage/logs/error.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
				if openErr == nil {
					if config.GetEnv().DEBUG {
						DefaultWriter = io.MultiWriter(file, os.Stdout)
					} else {
						DefaultWriter = io.MultiWriter(file)
					}

					fmt.Fprintf(DefaultWriter, "%s", "\n")
					//fmt.Fprintf(DefaultWriter, "%s %3d %s", red, "Error Msg: ", reset)
					fmt.Fprintf(DefaultWriter, "%s", "[" + time.Now().Format("2006-01-02 15:04:05") + "] app.ERROR: ")
					fmt.Fprintf(DefaultWriter, "%s", err)
					fmt.Fprintf(DefaultWriter, "%s", "\nStack trace:\n")
					fmt.Fprintf(DefaultWriter, "%s", debug.Stack())
					fmt.Fprintf(DefaultWriter, "%s", "\n")
				}

				c.JSON(200, gin.H{
					"code": 10500,
					"msg":  err,
				})
			}
		}()
		c.Next()
	}
}
