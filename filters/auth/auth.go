package auth

import (
	"github.com/gin-gonic/gin"
	"morningo/filters/auth/drivers"
	"net/http"
)

var driverList = map[string]Auth{
	"cookie": drivers.NewCookieAuthDriver(),
	"jwt":    drivers.NewJwtAuthDriver(),
}

type Auth interface {
	Check(c *gin.Context) bool
	User(c *gin.Context) interface{}
	Login(http *http.Request, w http.ResponseWriter, user map[string]interface{}) interface{}
	Logout(http *http.Request, w http.ResponseWriter) bool
}

func RegisterGlobalAuthDriver(authKey string, key string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(key, GenerateAuthDriver(authKey))
		c.Next()
	}
}

func Middleware(authKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !GenerateAuthDriver(authKey).Check(c) {
			c.HTML(http.StatusOK, "index.tpl", gin.H{
				"title": "login first",
			})
			c.Abort()
		}
		c.Next()
	}
}

func GenerateAuthDriver(string string) Auth {
	return driverList[string]
}

func GetCurUser(c *gin.Context, key string) map[string]interface{} {
	authDriver, _ := c.MustGet(key).(Auth)
	return authDriver.User(c).(map[string]interface{})
}

func User(c *gin.Context) map[string]interface{} {
	return GetCurUser(c, "jwt_auth")
}
