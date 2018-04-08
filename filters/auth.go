package filters

import (
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("X-Auth-Token")
		appid := c.Request.Header.Get("X-Auth-Appid")

		if token == "" && appid == "" {
			c.Abort()
			return
		}

		c.Next()
	}
}
