package controllers

import (
	db "gin-template/connections/database/mysql"
	"gin-template/module/cache"
	"gin-template/module/session"
	m "gin-template/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
	"fmt"
)

/**
 * 数据库与redis连接使用示例
 */
func IndexApi(c *gin.Context) {

	/** ORM使用 **/

	// Create
	m.Model.Create(&m.User{Name: "L1212", Avatar: "unknown", Sex: 1})

	// Read
	var user m.User
	m.Model.First(&user, 1) // find user with id 1
	m.Model.First(&user, "name = ?", "L1212") // find user with name l1212

	// Update
	m.Model.Model(&user).Update("avatar", "123")

	// Delete
	// m.Model.Delete(&user)

	/** 数据库连接使用 **/

	// 数据库插入
	insertRs := db.Exec("insert into users (name, avatar, sex) values (?, ?, ?)", "人才", "unknown", 1)
	insertId, _ := insertRs.LastInsertId()
	fmt.Printf("insert id: %d\n", insertId)

	// 数据库更新
	db.Exec("update users set name = ? where id = ?", "饭桶", insertId)

	// 数据库查询
	rs, con := db.Query("select name,avatar,id from users where id < ?", 100)
	defer con.Close() // 函数结束时关闭数据库连接

	/** 存储连接使用 **/

	// session存储
	session.GetStore().Set("key", "value", 0) // 0表示不过期
	str, _ := session.GetStore().Get("key")
	log.Println("session key: " + str)

	// cache存储
	cache.GetStore().Set("key", "value", time.Minute)
	cache.GetStore().Del("key")

	/** 数据返回 **/

	// 返回json数据
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": gin.H{
			"str": rs,
		},
	})

	// 返回字符串
	//c.String(http.StatusOK, "It works")

	// 返回html
	//c.HTML(http.StatusOK, "index.tpl", gin.H{
	//	"title": "Main website",
	//})
}
