package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"morningo/filters/auth/drivers"
)

var driverList = map[string]func()Auth {
	"cookie" : func() Auth {
		return drivers.NewCookieAuthDriver()
	},
	"jwt" : func() Auth {
		return drivers.NewJwtAuthDriver()
	},
}

type Auth interface {
	Check(c *gin.Context) bool
	User(c *gin.Context) interface{}
	Login(http *http.Request, w http.ResponseWriter, user map[string]interface{}) interface{}
	Logout(http *http.Request, w http.ResponseWriter) bool
}

func RegisterGlobalAuthDriver(authKey string, key string) gin.HandlerFunc {
	return func(c *gin.Context) {
		driver := GenerateAuthDriver(authKey)
		c.Set(key, driver)
		c.Next()
	}
}

func Middleware(authKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		driver := GenerateAuthDriver(authKey)
		if !(*driver).Check(c) {
			c.HTML(http.StatusOK, "index.tpl", gin.H{
				"title": "尚未登录，请登录",
			})
			c.Abort()
		}
		c.Next()
	}
}

func GenerateAuthDriver(string string) *Auth {
	var authDriver Auth
	authDriver = driverList[string]()
	return &authDriver
}