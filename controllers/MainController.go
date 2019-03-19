package controllers

import (
	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
	db "morningo/connections/database/mysql"
	"morningo/filters/auth"
	m "morningo/models"
	"net/http"
	"time"
	"morningo/connections/database"
)

func IndexApi(c *gin.Context) {

	// 返回html
	c.HTML(http.StatusOK, "index.tpl", gin.H{
		"title": "GO GO GO!",
	})
}

func DBexample(c *gin.Context) {

	// 数据库插入
	insertRs, _ := db.Exec("insert into users (name, avatar, sex) values (?, ?, ?)", "人才", "unknown", 1)
	insertId, _ := insertRs.LastInsertId()
	log.Printf("insert id: %d\n", insertId)

	// 数据库更新
	db.Exec("update users set name = ? where id = ?", "饭桶", insertId)

	// 数据库中间件
	database.Table("users").Where("id", "=", insertId).Update(database.H{
		"name": "你好",
	})

	// 数据库查询
	rs := db.Query("select name,avatar,id from users where id < ?", 100)
	log.Println(rs[0]["name"])

	rs1, _ := database.Table("users").
		Select("name", "avatar", "id").
		Where("id", "<", 100).
		All()
	log.Println(rs1[0])

	// 数据库事务
	db.WithTransaction(func(tx *db.SqlTxStruct) (error, map[string]interface{}) {
		_, err := tx.Query("select name,avatar,id from users where id < ?", 100)
		if err != nil {
			return err, map[string]interface{}{}
		}
		return nil, map[string]interface{}{}
	})

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
	session := sessions.Default(c)
	session.Set("key", "value") // 0表示不过期

	str := session.Get("key")
	log.Printf("session key: %s", str)

	// cache存储
	cacheStore, _ := c.MustGet(cache.CACHE_MIDDLEWARE_KEY).(*persistence.CacheStore)
	(*cacheStore).Set("key", "value", time.Minute)
	(*cacheStore).Delete("key")

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
	//log.Printf("user model insert %d\n", user.Model.ID)
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

func CookieSetExample(c *gin.Context) {
	authDr, _ := c.MustGet("web_auth").(*auth.Auth)

	id := c.Param("userid")

	rs := db.Query("select name,avatar,id from users where id = ?", id)

	log.Printf("len(rs): %d", len(rs))
	if len(rs) == 0 {
		c.HTML(http.StatusOK, "index.tpl", gin.H{
			"title": "wrong user id",
		})
		return
	}

	(*authDr).Login(c.Request, c.Writer, map[string]interface{}{"id": id})

	// 返回html
	c.HTML(http.StatusOK, "index.tpl", gin.H{
		"title": "login success!",
	})
}

func CookieGetExample(c *gin.Context) {
	authDr, _ := c.MustGet("web_auth").(*auth.Auth)

	userInfo := (*authDr).User(c).(map[interface{}]interface{})
	id, _ := userInfo["id"].(string)
	log.Println("id: " + id)

	rs := db.Query("select name,avatar,id from users where id = ?", id)

	// 返回html
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": gin.H{
			"user": rs,
		},
	})
}

func JwtSetExample(c *gin.Context) {
	authDr, _ := c.MustGet("jwt_auth").(*auth.Auth)

	token, _ := (*authDr).Login(c.Request, c.Writer, map[string]interface{}{"id": "123"}).(string)

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": gin.H{
			"token": token,
		},
	})
}

func JwtGetExample(c *gin.Context) {
	authDr, _ := c.MustGet("jwt_auth").(*auth.Auth)

	info := (*authDr).User(c).(map[string]interface{})

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": gin.H{
			"id": info["id"],
		},
	})
}
