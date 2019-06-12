package filters

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/contrib/sessions"
	"morningo/config"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-contrib/cache"
	"time"
)

func RegisterSession() gin.HandlerFunc {
	store, _ := sessions.NewRedisStore(
		10,
		"tcp",
		config.GetEnv().RedisIp+":"+config.GetEnv().RedisPort,
		config.GetEnv().RedisPassword,
		[]byte("secret"))
	return sessions.Sessions("mysession", store)
}

func RegisterCache() gin.HandlerFunc {
	var cacheStore persistence.CacheStore
	cacheStore = persistence.NewRedisCache(config.GetEnv().RedisIp+":"+config.GetEnv().RedisPort, config.GetEnv().RedisPassword, time.Minute)
	return cache.Cache(&cacheStore)
}
