package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Auth interface {
	Check(http *http.Request) bool
	User(http *http.Request) map[interface{}]interface{}
	Login(http *http.Request, w http.ResponseWriter, user map[interface{}]interface{}) bool
	Logout() bool
}

func AuthSetMiddleware(auth *Auth, key string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(key, auth)
		c.Next()
	}
}


func AuthMiddleware(auth *Auth) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !(*auth).Check(c.Request) {
			c.HTML(http.StatusOK, "index.tpl", gin.H{
				"title": "尚未登录，请登录",
			})
			c.Abort()
		}
		c.Next()
	}
}