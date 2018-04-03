package controllers

import (
	db "../connections/database/mysql"
	"../module/cache"
	"../module/session"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

/**
 * 数据库与redis连接使用示例
 */
func IndexApi(c *gin.Context) {
	rs, con := db.Query("select amount,name,avatar,id from users where id < ?", 20)
	defer con.Close() // 函数结束时关闭数据库连接

	// 0表示不过期
	session.GetStore().Set("key", "value", 0)
	str, _ := session.GetStore().Get("key")

	log.Println("session key: " + str)
	cache.GetStore().Set("key", "value", time.Minute)

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": gin.H{
			"str": rs,
		},
	})
	//c.String(http.StatusOK, "It works")
}
