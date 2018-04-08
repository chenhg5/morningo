package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	db "morningo/connections/database/mysql"
	m "morningo/models"
	"morningo/module/cache"
	"morningo/module/session"
	"net/http"
	"time"
)

func IndexApi(c *gin.Context) {

	// 返回html
	c.HTML(http.StatusOK, "index.tpl", gin.H{
		"title": "GO GO GO!",
	})
}

func DBexample(c *gin.Context) {

	// 数据库插入
	insertRs := db.Exec("insert into users (name, avatar, sex) values (?, ?, ?)", "人才", "unknown", 1)
	insertId, _ := insertRs.LastInsertId()
	fmt.Printf("insert id: %d\n", insertId)

	// 数据库更新
	db.Exec("update users set name = ? where id = ?", "饭桶", insertId)

	// 数据库查询
	rs, con := db.Query("select name,avatar,id from users where id < ?", 100)
	defer con.Close() // 函数结束时关闭数据库连接

	// 数据库事务
	Tx := db.BeginTransactions()
	_, err := Tx.Query("select name,avatar,id from users where id < ?", 100)
	if err != nil {
		Tx.Tx.Rollback()
	}
	Tx.Tx.Commit()

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": gin.H{
			"query_result": rs,
		},
	})
}

func StoreExample(c *gin.Context) {

	// session存储
	session.GetStore().Set("key", "value", 0) // 0表示不过期
	str, _ := session.GetStore().Get("key")
	log.Println("session key: " + str)

	// cache存储
	cache.GetStore().Set("key", "value", time.Minute)
	cache.GetStore().Del("key")

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": gin.H{
			"store_result": str,
		},
	})
}

func OrmExample(c *gin.Context) {

	// Create
	m.Model.Create(&m.User{Name: "L1212", Avatar: "unknown", Sex: 1})

	// Read
	var user m.User
	m.Model.First(&user, 1) // find user with id 1
	//fmt.Printf("user model insert %d\n", user.Model.ID)
	// m.Model.First(&user, "name = ?", "L1212") // find user with name l1212

	// Update
	m.Model.Model(&user).Update("avatar", "123456")

	// Delete
	// m.Model.Delete(&user)

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": gin.H{
			"orm_result": user,
		},
	})
}
