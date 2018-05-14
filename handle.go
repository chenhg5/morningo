package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"morningo/config"
	"os"
	"runtime/debug"
	"time"
	"net/http"
	"github.com/go-sql-driver/mysql"
)

var (
	defaultWriter io.Writer
)

func handleErrors() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {

				file, openErr := os.OpenFile(config.GetEnv().ERROR_LOG_PATH, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
				if openErr == nil {
					if config.GetEnv().DEBUG {
						defaultWriter = io.MultiWriter(file, os.Stdout)
					} else {
						defaultWriter = io.MultiWriter(file)
					}

					fmt.Fprintf(defaultWriter, "%s", "\n")
					//fmt.Fprintf(defaultWriter, "%s %3d %s", red, "Error Msg: ", reset)
					fmt.Fprintf(defaultWriter, "%s", "["+time.Now().Format("2006-01-02 15:04:05")+"] app.ERROR: ")
					fmt.Fprintf(defaultWriter, "%s", err)
					fmt.Fprintf(defaultWriter, "%s", "\nStack trace:\n")
					fmt.Fprintf(defaultWriter, "%s", debug.Stack())
					fmt.Fprintf(defaultWriter, "%s", "\n")
				}

				var (
					errMsg string
					mysqlError *mysql.MySQLError
					ok bool
				)
				if errMsg, ok = err.(string); ok {
					c.JSON(http.StatusInternalServerError, gin.H{
						"code": 500,
						"msg": "系统错误，错误代码为10000，错误信息：" + errMsg,
					})
					return
				} else if mysqlError, ok = err.(*mysql.MySQLError); ok {
					c.JSON(http.StatusInternalServerError, gin.H{
						"code": 500,
						"msg": "系统错误，错误代码为10000，错误信息：" + mysqlError.Error(),
					})
					return
				} else {
					c.JSON(http.StatusInternalServerError, gin.H{
						"code": 500,
						"msg": "系统错误，错误代码为10000",
					})
					return
				}
			}
		}()
		c.Next()
	}
}
