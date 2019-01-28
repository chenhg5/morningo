package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/go-sql-driver/mysql"
	"morningo/modules/log"
)

func handleErrors() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {

				log.Error(err)

				var (
					errMsg string
					mysqlError *mysql.MySQLError
					ok bool
				)
				if errMsg, ok = err.(string); ok {
					c.JSON(http.StatusInternalServerError, gin.H{
						"code": 500,
						"msg": "system error, " + errMsg,
					})
					return
				} else if mysqlError, ok = err.(*mysql.MySQLError); ok {
					c.JSON(http.StatusInternalServerError, gin.H{
						"code": 500,
						"msg": "system error, " + mysqlError.Error(),
					})
					return
				} else {
					c.JSON(http.StatusInternalServerError, gin.H{
						"code": 500,
						"msg": "system error",
					})
					return
				}
			}
		}()
		c.Next()
	}
}
